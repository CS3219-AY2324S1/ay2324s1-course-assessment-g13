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
	path.GITHUB_LOGIN: true,
	path.REFRESH:      true,
	"/":               true,
	"/github":         true,
}

func RequireAuthenticationMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := strings.Split(c.Request().RequestURI, "?")[0]
		_, isInList := bypassLoginList[uri]
		if isInList {
			return next(c)
		}

		tokenString, statusCode, responseMessage := cookie.Service.GetCookieValue(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		tokenClaims, statusCode, responseMessage := token.AccessTokenService.Validate(tokenString)
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		c.Set(TOKEN_CLAIMS_CONTEXT_KEY, tokenClaims)
		c.Request().Header.Set(USER_ROLE_KEY_REQUEST_HEADER, tokenClaims.User.Role)
		return next(c)
	}
}
