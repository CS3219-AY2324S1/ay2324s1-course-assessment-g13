package main

import (
	"question-service/config"
	"question-service/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	e := echo.New()

	questionGroup := e.Group("/questions")
	questionGroup.GET("", controllers.GetQuestions)
	questionGroup.GET("/:id", controllers.GetQuestion)
	questionGroup.POST("", controllers.CreateQuestion)
	questionGroup.DELETE("/:id", controllers.DeleteQuestion)

	e.Start(":8080")
}
