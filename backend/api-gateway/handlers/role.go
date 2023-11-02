package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/message"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const SUPER_ADMIN_ENV_KEY = "SUPER_ADMIN_KEY"

func UpgradeSuperAdmin(c echo.Context) error {
	superAdminKey := c.QueryParam("key")
	if superAdminKey != os.Getenv(SUPER_ADMIN_ENV_KEY) {
		return c.JSON(http.StatusForbidden, message.CreateErrorMessage(INVALID_SUPERADMIN_KEY))
	}

	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)
	userId := tokenClaims.User.ID
	var user *models.User
	err := config.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	currentRole := user.Role
	if currentRole == SUPER_ADMIN {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_ALREADY_SUPERADMIN))
	}

	user.Role = SUPER_ADMIN
	config.DB.Save(&user)

	c.Set(USER_CONTEXT_KEY, *user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_ROLE_UPGRADED_SUPER_ADMIN)

	return GenerateTokenAndSetCookie(c)
}

func UpgradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	superAdminId := tokenClaims.User.ID
	var superAdmin *models.User
	err := config.DB.Where("id = ?", superAdminId).First(&superAdmin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	currentSuperAdminRole := superAdmin.Role
	if currentSuperAdminRole != SUPER_ADMIN {
		return c.JSON(http.StatusForbidden, message.CreateErrorMessage(FAILURE_NOT_SUPERADMIN))
	}

	requestBody := new(models.LoginRequest)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var user models.User
	err = config.DB.Where("oauth_id = ? AND oauth_provider = ?", requestBody.OauthID, requestBody.OauthProvider).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	currentUserRole := user.Role
	if currentUserRole == ADMIN {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_HIGHEST))
	}

	user.Role = toggleUserRole(currentUserRole)
	config.DB.Save(&user)

	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_ROLE_UPGRADED))
}

func DowngradeUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	superAdminId := tokenClaims.User.ID
	var superAdmin *models.User
	err := config.DB.Where("id = ?", superAdminId).First(&superAdmin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	currentSuperAdminRole := superAdmin.Role
	if currentSuperAdminRole != SUPER_ADMIN {
		return c.JSON(http.StatusForbidden, message.CreateErrorMessage(FAILURE_NOT_SUPERADMIN))
	}

	requestBody := new(models.LoginRequest)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var user models.User
	err = config.DB.Where("oauth_id = ? AND oauth_provider = ?", requestBody.OauthID, requestBody.OauthProvider).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	currentUserRole := user.Role
	if currentUserRole == USER {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(FAILURE_USER_ROLE_LOWEST))
	}

	user.Role = toggleUserRole(currentUserRole)
	config.DB.Save(&user)

	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_ROLE_DOWNGRADED))
}

func toggleUserRole(role string) string {
	if role == USER {
		return ADMIN
	}
	return USER
}
