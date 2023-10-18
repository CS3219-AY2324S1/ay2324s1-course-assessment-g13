package handlers

import (
	"log"
	"net/http"

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
	config.DB.Where("user_id = ?", id).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, "User not found")
	}

	return c.String(http.StatusOK, user.Username)
}

func GetUsers(c echo.Context) error {
	users := make([]model.User, 0)
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user input")
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	payload := new(model.CreateUserPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var existingUser model.User
	err := config.DB.Where("auth_user_id = ?", payload.AuthUserID).First(&existingUser).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_EXIST))
	}

	var newUser model.User

	if err := config.DB.Where("username = ?", payload.Username).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusConflict, message.CreateErrorMessage(INVALID_USERNAME_EXIST))
	}
	newUser.Username = payload.Username

	newUser.AuthUserID = payload.AuthUserID
	newUser.PhotoUrl = payload.PhotoUrl
	newUser.PreferredLanguage = payload.PreferredLanguage
	err = config.DB.Create(&newUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessMessage(SUCCESS_USER_CREATED))
}

func UpdateUser(c echo.Context) error {
	authUserID := c.Param("authUserId")

	var user model.User
	if err := config.DB.Where("auth_user_id = ?", authUserID).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	payload := new(model.UpdateUserPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	if payload.PhotoUrl == "" && payload.Username == "" && payload.PreferredLanguage == "" {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_UPDATE_REQUEST))
	}

	if payload.Username != "" {
		var existingUser model.User
		if err := config.DB.Where("username = ?", payload.Username).First(&existingUser).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
			}
		} else {
			return c.JSON(http.StatusConflict, message.CreateErrorMessage(INVALID_USERNAME_EXIST))
		}
		log.Println(payload.Username)
		user.Username = payload.Username
	}

	if payload.PhotoUrl != "" {
		user.PhotoUrl = payload.PhotoUrl
	}

	if payload.PreferredLanguage != "" {
		user.PreferredLanguage = payload.PreferredLanguage
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_UPDATE_USER))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_UPDATED, user))

}

func DeleteUser(c echo.Context) error {
	authUserID := c.Param("authUserId")

	var existingUser model.User
	if err := config.DB.Where("auth_user_id = ?", authUserID).First(&existingUser).Error; err != nil {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	config.DB.Unscoped().Delete(&existingUser)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_USER_DELETED))
}
