package handlers

import (
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

var bypassLoginList = map[string]bool{
	"/auth/register": true,
	"/auth/login":    true,
}

func RequireAuthenticationMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := bypassLoginList[c.Request().RequestURI]
		if ok {
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

		c.Set(TOKEN_CLAIMS_KEY, tokenClaims)
		return next(c)
	}
}
