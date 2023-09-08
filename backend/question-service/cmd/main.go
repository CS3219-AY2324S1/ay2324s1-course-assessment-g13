package main

import (
	"fmt"
	"question-service/config"
	controller "question-service/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	fmt.Println("Starting development server")
	e := echo.New()

	questionGroup := e.Group("/questions")
	questionGroup.GET("", controller.GetQuestions)
	questionGroup.POST("", controller.CreateQuestion)

	e.Start(":8080")
}
