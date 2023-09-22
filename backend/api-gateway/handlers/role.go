package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpgradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_KEY).(*models.Claims)
	expirationTime := expiry.Add5MoreSeconds(tokenClaims.ExpiresAt.Time)

	userId := tokenClaims.User.ID
	var user *models.User
	config.DB.Where("id = ?", userId).First(&user)

	currentRole := user.Role
	if currentRole == ADMIN {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_HIGHEST))
	}

	user.Role = toggleUserRole(currentRole)
	config.DB.Save(&user)

	token, statusCode, responseMessage := token.Service.Generate(user, expirationTime)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	cookie_ := cookie.Service.CreateCookie(JWT_COOKIE_NAME, token, expirationTime)
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_ROLE_UPGRADED, *user))
}

func DowngradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_KEY).(*models.Claims)
	expirationTime := expiry.Add5MoreSeconds(tokenClaims.ExpiresAt.Time)

	userId := tokenClaims.User.ID
	var user *models.User
	config.DB.Where("id = ?", userId).First(&user)

	currentRole := user.Role
	if currentRole == USER {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_LOWEST))
	}

	user.Role = toggleUserRole(currentRole)
	config.DB.Save(&user)

	token, statusCode, responseMessage := token.Service.Generate(user, expirationTime)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	cookie_ := cookie.Service.CreateCookie(JWT_COOKIE_NAME, token, expirationTime)
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_ROLE_DOWNGRADED, *user))
}

func toggleUserRole(role string) string {
	if role == USER {
		return ADMIN
	}
	return USER
}
