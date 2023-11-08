package handlers

import (
	"fmt"
	"net/http"

	"user-service/config"
	model "user-service/models"

	"github.com/go-playground/validator/v10"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	req := new(model.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	// Validate the request input
	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user input")
	}

	var existingUser model.User
	config.DB.Where("username = ?", req.Username).First(&existingUser)
	if existingUser.ID != 0 {
		return c.JSON(http.StatusBadRequest, "Username already exists")
	}

	// Map CreateUserRequest fields to User model
	user := new(model.User)
	user.Username = req.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	user.HashedPassword = string(hashedPassword)

	user.UserID = req.UserID
	user.PhotoUrl = req.PhotoURL

	// Create a new user record in the database
	if err := config.DB.Create(user).Error; err != nil {
		errMsg := fmt.Sprintf("Failed to create user | err: %v", err)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	res := &model.LoginResponse{
		Id:       user.ID,
		Username: user.Username,
		PhotoUrl: user.PhotoUrl,
	}
	return c.JSON(http.StatusCreated, res)
}

func UpdateUserInfo(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := config.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "User not found")
	}

	req := new(model.UpdateUserInfo)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	updates := make(map[string]interface{})

	if req.Username != "" {
		updates["username"] = req.Username
	}

	if req.PhotoUrl != "" {
		updates["photo_url"] = req.PhotoUrl
	}

	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func UpdateUserPassword(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := config.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "User not found")
	}

	req := new(model.UpdateUserPassword)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.OldPassword)); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	if err := config.DB.Model(&user).Update("hashed_password", string(hashedPassword)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(http.StatusOK, "Password updated successfully")
}

func DeleteUser(c echo.Context) error {
	// sess, err := session.Get("session", c)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }

	// id := sess.Values["userId"]
	id := c.Param("id")

	var user model.User
	if err := config.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "User not found")
	}

	config.DB.Unscoped().Delete(&user)
	return c.JSON(http.StatusOK, "User deleted successfully")
}
