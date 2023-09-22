package main

import (
	"question-service/config"
	"question-service/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	config.PopulateDb()
	e := echo.New()

	questionGroup := e.Group("/questions")
	questionGroup.GET("", controllers.GetQuestions)
	questionGroup.GET("/:id", controllers.GetQuestion)
	questionGroup.POST("", controllers.CreateQuestion, controllers.AuthorizeAdminMiddleWare)
	questionGroup.DELETE("/:id", controllers.DeleteQuestion, controllers.AuthorizeAdminMiddleWare)
	questionGroup.PATCH("/:id", controllers.EditQuestion, controllers.AuthorizeAdminMiddleWare)

	e.Start(":8080")
}
