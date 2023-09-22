package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ADMIN                        = "admin"
	REQUEST_HEADER_USER_ROLE_KEY = "X-User-Role"
)

func AuthorizeAdminMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get(REQUEST_HEADER_USER_ROLE_KEY)
		if userRole == ADMIN {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Not Authorized!"})
	}
}
