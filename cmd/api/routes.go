package api

import (
	"broker/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

type errorresponse struct {
	Message string
}

func Routes(router *echo.Echo) {

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	router.POST("/ping", controller.Broker)
	router.POST("/auth/login", controller.Login)
	router.POST("/auth/register", controller.Register)

	router.Use(middlewareOne)
	router.GET("/auth/profile", controller.Login)
}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Register public route
		if c.Request().RequestURI == "/auth/login" {
			return next(c)
		}

		if c.Request().RequestURI == "/auth/register" {
			return next(c)
		}

		if c.Request().Header.Get("Authorization") == "" {
			response := errorresponse{
				Message: "Missing JWT",
			}

			return c.JSON(http.StatusOK, response)
		}

		// TODO : verif auth token

		return next(c)
	}
}
