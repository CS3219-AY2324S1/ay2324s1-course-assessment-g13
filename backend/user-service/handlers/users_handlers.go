package handlers

import (
	"net/http"

	"user-service/common/auth"
	"user-service/common/constants"
	"user-service/common/cookie"
	"user-service/common/errors"
	"user-service/config"
	model "user-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	config.DB.Where("id = ?", id).First(&user)
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

	// Create a new user record in the database
	if err := config.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

func UpdateUser(c echo.Context) error {
	claims := c.Get(constants.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)

	req := new(model.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.HashedPassword = string(hashedPassword)
	}

	config.DB.Save(&user)
	return c.JSON(http.StatusOK, "User updated successfully")
}

func DeleteUser(c echo.Context) error {
	claims := c.Get(constants.CLAIMS_KEY).(*model.Claims)
	userId := claims.User.ID

	var user model.User
	config.DB.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, "User not found")
	}

	config.DB.Delete(&user)
	cookie, err := cookie.SetCookieExpires(c.Cookie(constants.JWT_COOKIE_NAME))
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, message)
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "User deleted successfully")
}

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

	return c.JSON(http.StatusOK, "User Role Upgraded Successfully")
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

	return c.JSON(http.StatusOK, "User Role Downgraded Successfully")
}

func toggleRoles(currentRole string, isUpgrade bool) (string, error) {
	if currentRole == constants.BASIC_ROLE && !isUpgrade {
		return "", errors.MethodNotAllowedError("User has Lowest Role")
	}

	if currentRole == constants.ADMIN_ROLE && isUpgrade {
		return "", errors.MethodNotAllowedError("User has Highest Role")
	}

	if currentRole == constants.BASIC_ROLE {
		return constants.ADMIN_ROLE, nil
	}

	return constants.BASIC_ROLE, nil
}
