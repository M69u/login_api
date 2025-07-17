package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"login_rest_api/config"
	"login_rest_api/handlers"
)


func main() {

	// Initialize MongoDB connection
	config.ConnectDB()

	// Initialize Echo
	e := echo.New()

	// Middleware (optional but useful)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define routes
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

