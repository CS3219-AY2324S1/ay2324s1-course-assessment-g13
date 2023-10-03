package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpgradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	userId := tokenClaims.User.ID
	var user *models.User
	err := config.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	currentRole := user.Role
	if currentRole == ADMIN {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_HIGHEST))
	}

	user.Role = toggleUserRole(currentRole)
	config.DB.Save(&user)

	c.Set(USER_CONTEXT_KEY, *user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_ROLE_UPGRADED)

	return GenerateTokenAndSetCookie(c)
}

func DowngradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	userId := tokenClaims.User.ID
	var user *models.User
	err := config.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	currentRole := user.Role
	if currentRole == USER {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_LOWEST))
	}

	user.Role = toggleUserRole(currentRole)
	config.DB.Save(&user)

	c.Set(USER_CONTEXT_KEY, *user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_ROLE_DOWNGRADED)

	return GenerateTokenAndSetCookie(c)
}

func toggleUserRole(role string) string {
	if role == USER {
		return ADMIN
	}
	return USER
}
