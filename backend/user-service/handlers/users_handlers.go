package handlers

import (
	"net/http"
    "fmt"

	"user-service/models"
	"user-service/config"

	"github.com/labstack/echo/v4"
    "github.com/go-playground/validator/v10"
)

func GetUser(c echo.Context) error {
    // User ID from path `users/:id`
    id := c.Param("id")

    return c.String(http.StatusOK, "bye" + id)
}

func CreateUser(c echo.Context) error {
	req := new(models.CreateUserRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid JSON request")
    }

    // Validate the request input
    validator := validator.New()
    if err := validator.Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user input")
    }

    // Map CreateUserRequest fields to User model
    user := models.User{
        Username: req.Username,
        HashedPassword: req.Password,
    }

    fmt.Println("hello")

	// Create a new user record in the database
    if err := config.DB.Create(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create user")
    }

	return c.JSON(http.StatusCreated, "User created successfully")
}
