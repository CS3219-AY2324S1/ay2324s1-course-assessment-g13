package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const ADMIN = "admin"

func AllowAdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get("Authorization")
		if userRole != ADMIN {
			return c.JSON(http.StatusUnauthorized, "You are not authorized to perform this action")
		}
		return next(c)
	}
}
