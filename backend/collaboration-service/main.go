package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"log"
	"collaboration-service/handlers"
)

func main() {
	e := echo.New()

	hub := handlers.NewHub()
	wsHandler := handlers.NewHandler(hub)
	go hub.Run()

	e.GET("/ping", func(c echo.Context) error {
		log.Println("Receive request")
		return c.String(http.StatusOK, "I am the collaboration microservice!")
	})
	e.POST("/room", wsHandler.CreateRoom)
	e.GET("/ws/:roomId/:username", wsHandler.JoinRoom)
	e.GET("/ws/:roomId", wsHandler.GetQuestionId)
	e.Logger.Fatal(e.Start(":8080"))
}
