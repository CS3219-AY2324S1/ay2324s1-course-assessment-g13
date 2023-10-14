package handlers

import (
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	cookie_, statusCode, responseMessage = cookie.Service.SetCookieExpires(c.Cookie(REFRESH_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_LOGOUT))
}

func Refresh(c echo.Context) error {
	tokenString, statusCode, responseMessage := cookie.Service.GetCookieValue(c.Cookie(REFRESH_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	tokenClaims, statusCode, responseMessage := token.RefreshTokenService.Validate(tokenString)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	user := tokenClaims.User

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_TOKEN_REFRESHED)

	return GenerateTokenAndSetCookie(c)
}
