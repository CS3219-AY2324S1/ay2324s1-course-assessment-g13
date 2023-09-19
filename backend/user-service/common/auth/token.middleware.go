package auth

import (
	"net/http"
	"user-service/common/cookie"
	"user-service/common/errors"
	constants "user-service/common/utils"

	"github.com/labstack/echo/v4"
)

var noLoginRequiredList = map[string]bool{
	"/register": true,
	"/login":    true,
}

const USER_ALREADY_LOGGED_IN_MESSAGE = "You have already logged in"

func UserLoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := noLoginRequiredList[c.Request().RequestURI]
		if ok {
			return next(c)
		}

		tokenString, err := cookie.GetCookieValue(c.Cookie(constants.JWT_COOKIE_NAME))
		if err != nil {
			status, message := errors.ParseErrorToServiceError(err)
			return c.JSON(status, map[string]string{"message": message})
		}

		claims, err := TokenService.Validate(tokenString)
		if err != nil {
			status, message := errors.ParseErrorToServiceError(err)
			return c.JSON(status, map[string]string{"message": message})
		}

		c.Set(constants.CLAIMS_KEY, claims)

		return next(c)

	}
}

func PreventLoggedInUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionCookie, err := c.Cookie(constants.JWT_COOKIE_NAME)
		if err == nil && sessionCookie != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"message": USER_ALREADY_LOGGED_IN_MESSAGE})
		}
		return next(c)
	}
}
