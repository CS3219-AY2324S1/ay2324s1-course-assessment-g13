package controller

import (
	"context"
	"net/http"
	"question-service/config"
	"question-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetQuestions(c echo.Context) error {
	return c.String(http.StatusOK, "/questions")
}

func CreateQuestion(c echo.Context) error {
	var question models.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind request data"})
	}

	validator := validator.New()

	if err := validator.Struct(question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	collection := config.Collection
	existingQuestion := models.Question{}
	err := collection.FindOne(context.TODO(), bson.M{"title": question.Title}).Decode(&existingQuestion)
	if err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Question with this title already exists"})
	}

	result, err := collection.InsertOne(context.TODO(), question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert question"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Question created successfully", "insertedID": result.InsertedID})
}
