package handlers

import (
	"net/http"
	"strconv"

	"user-service/config"
	model "user-service/models"
	"user-service/utils/message"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := config.DB.Where("auth_user_id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_FOUND, user))
}

func GetUsers(c echo.Context) error {
	users := make([]model.User, 0)
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}
	return c.JSON(http.StatusOK, message.CreateSuccessUsersMessage(SUCCESS_USER_FOUND, users))
}

func CreateUser(c echo.Context) error {
	requestBody := new(model.CreateUser)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var existingUser model.User
	err := config.DB.Where("auth_user_id = ?", requestBody.AuthUserID).First(&existingUser).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_EXIST))
	}

	var newUser model.User

	if err := config.DB.Where("username = ?", requestBody.Username).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusConflict, message.CreateErrorMessage(INVALID_USERNAME_EXIST))
	}
	newUser.Username = requestBody.Username

	newUser.AuthUserID = requestBody.AuthUserID
	newUser.PhotoUrl = requestBody.PhotoUrl
	newUser.PreferredLanguage = requestBody.PreferredLanguage
	err = config.DB.Create(&newUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessMessage(SUCCESS_USER_CREATED))
}

func UpdateUser(c echo.Context) error {
	authUserID := c.Param("id")

	intAuthUserId, err := strconv.Atoi(authUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage("Internal Server Error: Unable to Convert String to Uint"))
	}

	uintAuthUserId := uint(intAuthUserId)

	var user model.User
	if err := config.DB.Where("auth_user_id = ?", authUserID).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	requestBody := new(model.UpdateUser)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	if requestBody.PhotoUrl == "" && requestBody.Username == "" && requestBody.PreferredLanguage == "" {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_UPDATE_REQUEST))
	}

	if requestBody.Username != "" {
		var existingUser model.User
		if err := config.DB.Where("username = ?", requestBody.Username).First(&existingUser).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
			}
		} else {
			if existingUser.AuthUserID != uintAuthUserId {
				return c.JSON(http.StatusConflict, message.CreateErrorMessage(INVALID_USERNAME_EXIST))
			}
		}
		user.Username = requestBody.Username
	}

	if requestBody.PhotoUrl != "" {
		user.PhotoUrl = requestBody.PhotoUrl
	}

	if requestBody.PreferredLanguage != "" {
		user.PreferredLanguage = requestBody.PreferredLanguage
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_UPDATE_USER))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_UPDATED, user))

}

func DeleteUser(c echo.Context) error {
	authUserID := c.Param("id")

	var existingUser model.User
	if err := config.DB.Where("auth_user_id = ?", authUserID).First(&existingUser).Error; err != nil {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	config.DB.Unscoped().Delete(&existingUser)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_USER_DELETED))
}
