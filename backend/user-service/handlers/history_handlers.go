package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"user-service/config"
	model "user-service/models"
	"user-service/utils/message"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateHistory(c echo.Context) error {
	reqBody := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err := config.DB.Where("username = ?", reqBody["username"]).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	var language string
	if reqBody["language"] == "" {
		language = user.PreferredLanguage
	} else {
		language = reqBody["language"]
	}

	var existingHistory model.History
	err = config.DB.Where("user_id = ? AND room_id = ? AND question_id = ?", user.ID, reqBody["room_id"], reqBody["question_id"]).First(&existingHistory).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		// Update existing history
		existingHistory.Solution = reqBody["solution"]
		existingHistory.Language = language
		if err := config.DB.Save(&existingHistory).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}

		return c.JSON(http.StatusCreated, message.CreateSuccessMessage("History updated"))
	}

	// Create new history
	var newHistory model.History
	newHistory.UserID = user.ID
	newHistory.RoomId = reqBody["room_id"]
	newHistory.QuestionId = reqBody["question_id"]
	newHistory.Title = reqBody["title"]
	newHistory.Solution = reqBody["solution"]
	newHistory.Language = language

	if err := config.DB.Create(&newHistory).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessMessage("History created"))
}	

func GetHistories(c echo.Context) error {
	authId := c.Param("authId")
	user := model.User{}

	if err := config.DB.Where("auth_user_id = ?", authId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		log.Println("db error")
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	histories := make([]model.History, 0)
	if err := config.DB.Where("user_id = ?", user.ID).Order("updated_at DESC").Find(&histories).Error; err != nil {

		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessHistoriesMessage(SUCCESS_HISTORY_FOUND, histories))
}

