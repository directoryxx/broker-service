package http

import (
	"broker/cmd/api"
	"broker/infrastructure"
	"broker/internal/domain"
	"broker/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func RunWebserver() {
	app := echo.New()

	app.Validator = &CustomValidator{validator: validator.New()}

	log.Println("[INFO] Starting Broker Service on port", os.Getenv("APPLICATION_PORT"))

	log.Println("[INFO] Loading Redis")
	rediisConn := infrastructure.OpenRedis()

	log.Println("[INFO] Loading Kafka Producer")
	kafkaProducer, err := infrastructure.ConnectKafka()

	if err != nil {
		log.Fatalf("Could not initialize connection to kafka producer %s", err)
	}

	log.Println("[INFO] Loading Repository")
	userRepo := repository.NewUserRepository(rediisConn, kafkaProducer)

	log.Println("[INFO] Loading Middleware")
	SetMiddleware(app, userRepo)

	log.Println("[INFO] Loading Routes")
	api.Routes(app)

	log.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT"))))
}

func SetMiddleware(r *echo.Echo, userRepo repository.UserRepository) {
	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogMethod:    true,
		LogHost:      true,
		LogLatency:   true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			// log :=
			hostname, _ := os.Hostname()
			log := &domain.Log{
				RemoteIP:      values.RemoteIP,
				Service:       os.Getenv("APPLICATION_NAME") + " Service",
				ContainerName: hostname,
				Time:          values.StartTime.String(),
				Host:          values.Host,
				Method:        values.Method,
				Uri:           values.URI,
				UserAgent:     values.UserAgent,
				Status:        strconv.Itoa(values.Status),
				Latency:       values.Latency.String(),
				LatencyHuman:  values.Latency.String(),
				// Error:         values.Error.Error(),
			}

			b, _ := json.Marshal(log)

			userRepo.Publish(c.Request().Context(), string(b), "log-http")

			return nil
		},
	}))
	r.Use(middleware.Recover())
}
