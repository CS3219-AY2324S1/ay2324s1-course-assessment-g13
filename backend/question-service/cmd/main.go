package main

import (
	auth "question-service/common"
	"question-service/config"
	"question-service/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	config.PopulateDb()
	e := echo.New()

	questionGroup := e.Group("/questions")
	questionGroup.GET("/", controllers.GetQuestions)
	questionGroup.GET("/:id", controllers.GetQuestion)
	questionGroup.POST("/", controllers.CreateQuestion, auth.AllowAdminOnly)
	questionGroup.DELETE("/:id", controllers.DeleteQuestion, auth.AllowAdminOnly)
	questionGroup.PATCH("/:id", controllers.EditQuestion, auth.AllowAdminOnly)

	e.Start(":8080")
}
