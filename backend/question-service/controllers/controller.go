package controllers

import (
	"context"
	"net/http"
	"question-service/config"
	"question-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuestion(c echo.Context) error {
	questionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	question := models.Question{}
	err = config.Collection.FindOne(context.TODO(), bson.M{"_id": questionID}).Decode(&question)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Specified question does not exist"})
	}

	return c.JSON(http.StatusOK, question)
}

func GetQuestions(c echo.Context) error {
	cursor, err := config.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve questions"})
	}
	defer cursor.Close(context.TODO())

	var questions []models.Question
	if err := cursor.All(context.TODO(), &questions); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode questions"})
	}

	return c.JSON(http.StatusOK, questions)
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

	err := config.Collection.FindOne(context.TODO(), bson.M{"title": question.Title}).Err()
	if err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Question with this title already exists"})
	}

	result, err := config.Collection.InsertOne(context.TODO(), question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert question"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Question created successfully", "insertedID": result.InsertedID})
}

func DeleteQuestion(c echo.Context) error {
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	result, err := config.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Specified question does not exist"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Question deleted successfully"})
}

func EditQuestion(c echo.Context) error {
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	collection := config.Collection
	existingQuestion := models.Question{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&existingQuestion)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Specified question does not exist"})
	}

	var request models.EditRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to bind request data"})
	}

	validator := validator.New()
	if err := validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if request.Title != "" {
		filter := bson.M{
			"_id":   bson.M{"$ne": objectID},
			"title": request.Title,
		}
		if err = collection.FindOne(context.TODO(), filter).Err(); err == nil {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Another question with this title already exists"})
		}
	}

	if _, err = collection.UpdateByID(context.Background(), objectID, bson.M{"$set": request}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Question successfully updated"})
}
