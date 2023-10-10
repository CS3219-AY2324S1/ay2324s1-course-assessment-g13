package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_FOUND, user))
}

func DeleteUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	err := config.DB.Delete(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_DELETED, user))
}
