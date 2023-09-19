package handlers

import (
	"net/http"
	"user-service/common/auth"
	"user-service/common/cookie"
	"user-service/common/errors"
	constants "user-service/common/utils"
	"user-service/config"
	model "user-service/models"

	"github.com/labstack/echo/v4"
)

const (
	UPGRADE_SUCCESS_MESSAGE   = "User Role Upgraded Successfully"
	DOWNGRADE_SUCCESS_MESSAGE = "User Role Downgraded Successfully"
	UPGRADE_FAIL_MESSAGE      = "User has Highest Role"
	DOWNGRADE_FAIL_MESSAGE    = "User has Lowest Role"
)

func UpgradeRole(c echo.Context) error {
	claims := c.Get(constants.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID
	userRole := claims.User.Role

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)

	newRole, err := toggleRoles(userRole, true)
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}

	user.Role = newRole

	config.DB.Save(&user)

	expirationTime := auth.GetExpirationTime()

	newTokenString, err := auth.TokenService.Generate(&user, expirationTime)
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}

	cookie := cookie.CreateCookie(constants.JWT_COOKIE_NAME, newTokenString, expirationTime)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, UPGRADE_SUCCESS_MESSAGE)
}

func DowngradeRole(c echo.Context) error {
	claims := c.Get(constants.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID
	userRole := claims.User.Role

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)

	newRole, err := toggleRoles(userRole, false)
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}

	user.Role = newRole

	config.DB.Save(&user)

	expirationTime := auth.GetExpirationTime()

	newTokenString, err := auth.TokenService.Generate(&user, expirationTime)
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}

	cookie := cookie.CreateCookie(constants.JWT_COOKIE_NAME, newTokenString, expirationTime)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, DOWNGRADE_SUCCESS_MESSAGE)
}

func toggleRoles(currentRole string, isUpgrade bool) (string, error) {
	if currentRole == constants.BASIC_ROLE && !isUpgrade {
		return "", errors.MethodNotAllowedError(DOWNGRADE_FAIL_MESSAGE)
	}

	if currentRole == constants.ADMIN_ROLE && isUpgrade {
		return "", errors.MethodNotAllowedError(UPGRADE_FAIL_MESSAGE)
	}

	if currentRole == constants.BASIC_ROLE {
		return constants.ADMIN_ROLE, nil
	}

	return constants.BASIC_ROLE, nil
}
