package routes

import (
	"linkqu-be-go/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})

	e.POST("/users", handlers.CreateUser)
}
