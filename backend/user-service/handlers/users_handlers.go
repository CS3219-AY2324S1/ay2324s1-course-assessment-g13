package handlers

import (
	"net/http"

	"user-service/config"
	model "user-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, "bye"+id)
}

func GetUsers(c echo.Context) error {
	users := make([]model.User, 0)
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user input")
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	req := new(model.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	// Validate the request input
	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user input")
	}

	// Map CreateUserRequest fields to User model
	user := new(model.User)
	user.Username = req.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	user.HashedPassword = string(hashedPassword)

	// Create a new user record in the database
	if err := config.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}
