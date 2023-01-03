package controller

import (
	"broker/internal/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type brokerresponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type profileresponse struct {
	Error   bool         `json:"error"`
	Message string       `json:"message"`
	Profile *domain.User `json:"profile"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type errorresponse struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
}

type loginresponse struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
	Data    any  `json:"data"`
	Token   any  `json:"token"`
}

func Broker(c echo.Context) error {
	response := brokerresponse{
		Error:   false,
		Message: "Hit the broker",
	}

	return c.JSON(http.StatusOK, response)
}

func Login(c echo.Context) error {
	// Validation
	u := new(AuthPayload)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		fmt.Println(err)
		return err
	}

	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(u, "", "\t")

	// call the service
	// request, err := http.NewRequest("POST", "http://127.0.0.1:8016/login", bytes.NewBuffer(jsonData))
	request, err := http.NewRequest("POST", "http://auth-service/login", bytes.NewBuffer(jsonData))
	if err != nil {
		response := errorresponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	// add application/json to header
	request.Header.Add("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		response := errorresponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	defer response.Body.Close()

	// Pass response
	if response.Status == "200 OK" {
		loginRes := &loginresponse{}
		body, _ := io.ReadAll(response.Body)
		var jsonData = []byte(body)

		var _ = json.Unmarshal(jsonData, &loginRes)

		responseBroker := brokerresponse{
			Error:   false,
			Message: "Response Auth Service",
			Data:    loginRes,
		}

		return c.JSON(http.StatusOK, responseBroker)
	} else {
		loginRes := &loginresponse{}
		body, _ := io.ReadAll(response.Body)
		var jsonData = []byte(body)

		var _ = json.Unmarshal(jsonData, &loginRes)
		responseBroker := brokerresponse{
			Error:   false,
			Message: "Failed",
			Data:    loginRes,
		}

		return c.JSON(http.StatusOK, responseBroker)
	}
}

func Register(c echo.Context) error {
	// Validation
	u := new(RegisterPayload)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		fmt.Println(err)
		return err
	}

	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(u, "", "\t")

	// call the service
	// request, err := http.NewRequest("POST", "http://127.0.0.1:8016/login", bytes.NewBuffer(jsonData))
	request, err := http.NewRequest("POST", "http://auth-service/register", bytes.NewBuffer(jsonData))
	if err != nil {
		response := errorresponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	// add application/json to header
	request.Header.Add("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		response := errorresponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	defer response.Body.Close()

	// Pass response
	if response.Status == "200 OK" {
		loginRes := &loginresponse{}
		body, _ := io.ReadAll(response.Body)
		var jsonData = []byte(body)

		var _ = json.Unmarshal(jsonData, &loginRes)

		responseBroker := brokerresponse{
			Error:   false,
			Message: "Response Auth Service",
			Data:    loginRes,
		}

		return c.JSON(http.StatusOK, responseBroker)
	} else {
		loginRes := &loginresponse{}
		body, _ := io.ReadAll(response.Body)
		var jsonData = []byte(body)

		var _ = json.Unmarshal(jsonData, &loginRes)
		responseBroker := brokerresponse{
			Error:   false,
			Message: "Failed",
			Data:    loginRes,
		}

		return c.JSON(http.StatusOK, responseBroker)
	}
}

func Profile(c echo.Context) error {
	user := c.Get("user").(domain.User)
	responseBroker := &profileresponse{
		Error:   false,
		Message: "Berhasil mengambil data",
		Profile: &user,
	}
	return c.JSON(http.StatusOK, responseBroker)
}
