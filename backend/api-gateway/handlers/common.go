package handlers

import (
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GenerateTokenAndSetCookie(c echo.Context) error {
	user := c.Get(USER_CONTEXT_KEY).(models.User)
	successMessage := c.Get(SUCCESS_MESSAGE_CONTEXT_KEY).(string)
	expirationTime := c.Get(EXPIRATION_TIME_CONTEXT_KEY).(time.Time)

	tokenString, statusCode, responseMessage := token.Service.Generate(&user, expirationTime)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	cookie_ := cookie.Service.CreateCookie(JWT_COOKIE_NAME, tokenString, expirationTime)
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(successMessage, user))
}
