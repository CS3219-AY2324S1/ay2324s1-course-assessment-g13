package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"log"

	"github.com/labstack/echo/v4"
)

const (
	REGISTER  = "/auth/register"
	AUTH_USER = "/auth/user"
	LOGIN     = "/auth/login"
	LOGOUT    = "/auth/logout"
)

const API_GATEWAY_PORT = ":1234"

func main() {
	config.ConnectDb()
	log.Println("Starting development server...")

	API_GATEWAY := echo.New()

	API_GATEWAY.Use(handlers.RequireAuthenticationMiddleWare)

	API_GATEWAY.POST(REGISTER, handlers.CreateUser)
	API_GATEWAY.GET(AUTH_USER, handlers.GetUser)
	API_GATEWAY.DELETE(AUTH_USER, handlers.DeleteUser)
	API_GATEWAY.POST(LOGIN, handlers.Login)
	API_GATEWAY.GET(LOGOUT, handlers.Logout)

	API_GATEWAY.Start(API_GATEWAY_PORT)
}
