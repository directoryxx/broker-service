package api

import (
	"broker/internal/controller"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

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

}
