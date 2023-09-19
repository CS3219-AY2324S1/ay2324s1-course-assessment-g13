package handlers

import (
	"net/http"

	"user-service/common/cookie"
	"user-service/common/errors"
	"user-service/common/utils"
	"user-service/config"
	model "user-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER_NOT_FOUND_MESSAGE       = "User Not Found"
	INVALID_USER_INPUT_MESSAGE   = "Invalid User Input"
	USERNAME_EXIST_MESSAGE       = "Username Already Exists"
	USER_CREATED_SUCCESS_MESSAGE = "User Created Successfully"
	USER_CREATED_FAIL_MESSAGE    = "Failed to Create User"
	USER_UPDATED_SUCCESS_MESSAGE = "User Updated Successfully"
	USER_DELETED_SUCCESS_MESSAGE = "User Deleted Successfully"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	config.DB.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, USER_NOT_FOUND_MESSAGE)
	}

	return c.JSON(http.StatusOK, user.Username)
}

func GetUsers(c echo.Context) error {
	users := make([]model.User, 0)
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, INVALID_USER_INPUT_MESSAGE)
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	req := new(model.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, INVALID_JSON_REQUEST_MESSAGE)
	}

	// Validate the request input
	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, INVALID_USER_INPUT_MESSAGE)
	}

	var existingUser model.User
	config.DB.Where("username = ?", req.Username).First(&existingUser)
	if existingUser.ID != 0 {
		return c.JSON(http.StatusBadRequest, USERNAME_EXIST_MESSAGE)
	}

	// Map CreateUserRequest fields to User model
	user := new(model.User)
	user.Username = req.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.INTERNAL_SERVER_ERROR_MESSAGE)
	}
	user.HashedPassword = string(hashedPassword)

	// Create a new user record in the database
	if err := config.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, USER_CREATED_FAIL_MESSAGE)
	}

	return c.JSON(http.StatusCreated, USER_CREATED_SUCCESS_MESSAGE)
}

func UpdateUser(c echo.Context) error {
	claims := c.Get(utils.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)

	req := new(model.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, INVALID_JSON_REQUEST_MESSAGE)
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.HashedPassword = string(hashedPassword)
	}

	config.DB.Save(&user)
	return c.JSON(http.StatusOK, USER_UPDATED_SUCCESS_MESSAGE)
}

func DeleteUser(c echo.Context) error {
	claims := c.Get(utils.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, USER_NOT_FOUND_MESSAGE)
	}

	config.DB.Delete(&user)
	cookie, err := cookie.SetCookieExpires(c.Cookie(utils.JWT_COOKIE_NAME))
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, USER_DELETED_SUCCESS_MESSAGE)
}
