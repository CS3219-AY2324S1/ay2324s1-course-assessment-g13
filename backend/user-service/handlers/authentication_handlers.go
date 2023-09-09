package handlers

import (
	"net/http"
	"user-service/config"
	model "user-service/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	var user model.User
	config.DB.Where("username = ?", req.Username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		// Passwords don't match
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	return c.JSON(http.StatusOK, "Login successful")

}
