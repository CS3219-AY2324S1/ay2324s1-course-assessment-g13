package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/client"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	oauthId := tokenClaims.User.OauthID
	oauthProvider := tokenClaims.User.OauthProvider
	var superAdminUser models.User
	err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", oauthId, oauthProvider).First(&superAdminUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	if superAdminUser.Role != SUPER_ADMIN {
		return c.JSON(http.StatusForbidden, message.CreateErrorMessage(FAILURE_NOT_SUPERADMIN_GET_USERS))
	}

	users := make([]models.User, 0)
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}
	return c.JSON(http.StatusOK, message.CreateSuccessUsersMessage(SUCCESS_USER_FOUND, users))
}

func GetUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	currentUser := tokenClaims.User
	oauthId := currentUser.OauthID
	oauthProvider := currentUser.OauthProvider

	var user models.User
	err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", oauthId, oauthProvider).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_FOUND, user))
}

func DeleteUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	if user.Role == SUPER_ADMIN {
		return c.JSON(http.StatusForbidden, message.CreateErrorMessage(FAILURE_DELETE_SUPERADMIN))
	}

	authUserId := user.ID
	responseStatusCode, responseMessage := client.UserService.DeleteUser(authUserId)
	if responseStatusCode != http.StatusOK {
		return c.JSON(responseStatusCode, message.CreateErrorMessage(responseMessage))
	}

	err := config.DB.Delete(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	cookie_, statusCode, responseMessage = cookie.Service.SetCookieExpires(c.Cookie(REFRESH_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_DELETED, user))
}

func CreateUser(c echo.Context) error {
	requestBody := new(models.CreateUser)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var existingUser models.User
	err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", requestBody.OauthID, requestBody.OauthProvider).First(&existingUser).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_EXIST))
	}

	var newUser models.User
	newUser.OauthID = requestBody.OauthID
	newUser.OauthProvider = requestBody.OauthProvider
	err = config.DB.Create(&newUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}

	userServiceCreateUserRequestBody := models.UserServiceCreateUserRequestBody{
		AuthUserID:        newUser.ID,
		Username:          requestBody.Username,
		PhotoUrl:          requestBody.PhotoUrl,
		PreferredLanguage: requestBody.PreferredLanguage,
	}

	responseStatusCode, responseMessage := client.UserService.CreateUser(userServiceCreateUserRequestBody)

	if responseStatusCode != http.StatusCreated {
		err := config.DB.Delete(&newUser).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
		}
		return c.JSON(responseStatusCode, message.CreateErrorMessage(responseMessage))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessUserMessage(SUCCESS_USER_CREATED, newUser))
}

func UpdateUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	currentUser := tokenClaims.User
	authId := currentUser.ID

	requestBody := new(models.UpdateUser)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage("Error marshalling request body"))
	}

	resp, err := http.Post(os.Getenv("USER_SERVICE_URL") + "/users/" + strconv.FormatUint(uint64(authId), 10), "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage("Error creating request"))
	}
	defer resp.Body.Close()

	var updateUserResponse models.UpdateUserResponse
	err = json.NewDecoder(resp.Body).Decode(&updateUserResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}
	
	return c.JSON(http.StatusOK, message.CreateSuccessUupdateUserMessage(SUCCESS_USER_UPDATED, updateUserResponse.User))
}
