package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// Check if user exist
	// Create Token
	// Save Token as Cookie
	return c.JSON(http.StatusOK, "Login")
}

func Logout(c echo.Context) error {
	// Set Token to Expire in Cookie
	return c.JSON(http.StatusOK, "Logout")
}
