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
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	question := models.Question{}
	err = config.Collection.FindOne(context.TODO(), bson.M{"_id": questionID}).Decode(&question)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, "Specified question does not exist")
	}

	return c.JSON(http.StatusOK, question)
}

func GetRandomQuestionId(c echo.Context) error {
	complexity := c.Param("complexity")
	filter := bson.M{"complexity": complexity}
	pipeline := []bson.M{
		{"$match": filter},
		{"$sample": bson.M{"size": 1}},
	}
	cursor, err := config.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve questions")
	}
	defer cursor.Close(context.TODO())

	question := models.Question{}
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&question)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to decode questions")
		}
	}

	return c.JSON(http.StatusOK, question.Id)
}

func GetQuestions(c echo.Context) error {
	cursor, err := config.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve questions")
	}
	defer cursor.Close(context.TODO())

	var questions []models.Question
	if err := cursor.All(context.TODO(), &questions); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to decode questions")
	}

	return c.JSON(http.StatusOK, questions)
}

func GetCategories(c echo.Context) error {
	cursor, err := config.CategoriesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve categories")
	}
	defer cursor.Close(context.TODO())

	var categories []models.Category
	if err := cursor.All(context.TODO(), &categories); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to decode categories")
	}

	return c.JSON(http.StatusOK, categories)
}

func CreateQuestion(c echo.Context) error {
	var question models.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to bind request data")
	}

	validator := validator.New()
	if err := validator.Struct(question); err != nil {
		return c.JSON(http.StatusBadRequest, "Inputted data is invalid")
	}

	err := config.Collection.FindOne(context.TODO(), bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: "^" + question.Title + "$", Options: "i"}}}).Err()
	if err == nil {
		return c.JSON(http.StatusConflict, "Question with this title already exists")
	}

	result, err := config.Collection.InsertOne(context.TODO(), question)
	_ = result
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to insert question")
	}

	return c.JSON(http.StatusCreated, "Question created successfully")
}

func DeleteQuestion(c echo.Context) error {
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	result, err := config.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete question")
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, "Specified question does not exist")
	}

	return c.JSON(http.StatusOK, "Question deleted successfully")
}
