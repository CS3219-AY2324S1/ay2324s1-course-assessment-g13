package handlers

import (
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(JWT_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_LOGOUT))
}

func Refresh(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_TOKEN_REFRESHED)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expiry.ExpireIn5Minutes())

	return GenerateTokenAndSetCookie(c)
}
