package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpgradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)
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

	c.Set(USER_CONTEXT_KEY, *user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_ROLE_UPGRADED)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expirationTime)

	return GenerateTokenAndSetCookie(c)
}

func DowngradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)
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

	c.Set(USER_CONTEXT_KEY, *user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_ROLE_DOWNGRADED)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expirationTime)

	return GenerateTokenAndSetCookie(c)
}

func toggleUserRole(role string) string {
	if role == USER {
		return ADMIN
	}
	return USER
}
