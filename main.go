package main

import (
	"linkqu-be-go/config"
	"linkqu-be-go/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Bisa disesuaikan dengan domain frontend
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	config.ConnectDB()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
