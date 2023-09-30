package handlers

import (
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenerateTokenAndSetCookie(c echo.Context) error {
	user := c.Get(USER_CONTEXT_KEY).(models.User)
	successMessage := c.Get(SUCCESS_MESSAGE_CONTEXT_KEY).(string)
	accessTokenExpiry := expiry.ExpireIn5Minutes()
	refreshTokenExpiry := expiry.ExpireIn24Hours()

	accessTokenString, statusCode, responseMessage := token.AccessTokenService.Generate(&user, accessTokenExpiry)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	refreshTokenString, statusCode, responseMessage := token.RefreshTokenService.Generate(&user, refreshTokenExpiry)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	cookie_ := cookie.Service.CreateCookie(ACCESS_TOKEN_COOKIE_NAME, accessTokenString, accessTokenExpiry)
	c.SetCookie(cookie_)
	cookie_ = cookie.Service.CreateCookie(REFRESH_TOKEN_COOKIE_NAME, refreshTokenString, refreshTokenExpiry)
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(successMessage, user))
}
