package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"producer/models"
	"producer/utils"
)

func UserCancelHandler(c echo.Context) error {
	// Destructures the API request body into our model
	requestBody := models.CancelRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		msg := fmt.Sprintf("[UserCancelHandler] Error decoding request body | err: %v", err)
		log.Println(msg)
		return err
	}
	// Indicate user has cancelled
	utils.CancelUser(requestBody.Username)

	cancelResponseBody := models.CancelResponse{CancelStatus: true}
	return c.JSON(http.StatusOK, cancelResponseBody)
}
