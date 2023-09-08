package controller

import (
	"context"
	"net/http"
	"question-service/config"
	"question-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetQuestions(c echo.Context) error {
	return c.String(http.StatusOK, "/questions")
}

func CreateQuestion(c echo.Context) error {
	var question models.Question
	if err := c.Bind(&question); err != nil {
		return err
	}

	validator := validator.New()
	/*	if err := validator.RegisterValidation("validComplexity", validateComplexity); err != nil {
		log.Fatalf("Failed to register custom validator: %v", err)
	} */

	if err := validator.Struct(question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	collection := config.Collection
	result, err := collection.InsertOne(context.TODO(), question)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result.InsertedID)
}

/*
func validateComplexity(fl validator.FieldLevel) bool {
	complexity := fl.Field().String()
	fmt.Println(complexity)
	return complexity == "Easy" || complexity == "Medium" || complexity == "Hard"
} */
