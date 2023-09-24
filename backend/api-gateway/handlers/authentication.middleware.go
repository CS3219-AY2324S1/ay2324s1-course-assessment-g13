package handlers

import (
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/path"
	"api-gateway/utils/token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var bypassLoginList = map[string]bool{
	path.REGISTER:     true,
	path.LOGIN:        true,
	path.GITHUB_LOGIN: true,
}

func RequireAuthenticationMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := strings.Split(c.Request().RequestURI, "?")[0]
		_, isInList := bypassLoginList[uri]
		if isInList {
			return next(c)
		}

		tokenString, statusCode, responseMessage := cookie.Service.GetCookieValue(c.Cookie(JWT_COOKIE_NAME))
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		tokenClaims, statusCode, responseMessage := token.Service.Validate(tokenString)
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		c.Set(TOKEN_CLAIMS_CONTEXT_KEY, tokenClaims)
		c.Request().Header.Set(REQUEST_HEADER_USER_ROLE_KEY, tokenClaims.User.Role)
		return next(c)
	}
}

func PreventLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, isInList := bypassLoginList[c.Request().RequestURI]
		if !isInList {
			return next(c)
		}
		_, statusCode, _ := cookie.Service.GetCookieValue(c.Cookie(JWT_COOKIE_NAME))
		if statusCode == http.StatusOK {
			return c.JSON(http.StatusForbidden, message.CreateErrorMessage(FAILURE_USER_ALREADY_LOGIN))
		}
		return next(c)
	}
}
