package main

import (
	"linkqu-be-go/config"
	"linkqu-be-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.ConnectDB()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
