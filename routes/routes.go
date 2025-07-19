package routes

import (
	"login_rest_api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)
}
