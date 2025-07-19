package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"login_rest_api/config"
	"login_rest_api/routes"
)

func main() {

	// Initialize MongoDB connection
	config.ConnectDB()

	// Initialize Echo
	e := echo.New()

	// Middleware (optional but useful)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
