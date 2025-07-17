package handlers

import (
	"context"
	"net/http"
	"time"

	"login_rest_api/config"
	"login_rest_api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// RegisterUser handles user registration
func RegisterUser(c echo.Context) error {
	var user models.User

	// Bind request body to user struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	// Get MongoDB collection
	collection := config.DB.Collection("users")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists by phone
	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"phone": user.Phone}).Decode(&existingUser)
	if err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User already registered",
		})
	}

	// Insert new user
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to register user",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
	})
}

// LoginUser handles user login
func LoginUser(c echo.Context) error {
	var credentials models.User

	// Bind login request
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid login data",
		})
	}

	// Get MongoDB collection
	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user by phone and name
	var foundUser models.User
	err := collection.FindOne(ctx, bson.M{
		"phone": credentials.Phone,
		"name":  credentials.Name,
	}).Decode(&foundUser)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid phone or name",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
	})
}
