package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/utils/path"
	"log"

	"github.com/labstack/echo/v4"
)

const API_GATEWAY_PORT = ":1234"

func main() {
	config.ConnectDb()
	log.Println("Starting development server...")

	API_GATEWAY := echo.New()

	API_GATEWAY.Use(handlers.PreventLoginMiddleware, handlers.RequireAuthenticationMiddleWare)

	API_GATEWAY.POST(path.REGISTER, handlers.CreateUser)
	API_GATEWAY.GET(path.AUTH_USER, handlers.GetUser)
	API_GATEWAY.DELETE(path.AUTH_USER, handlers.DeleteUser)

	API_GATEWAY.POST(path.LOGIN, handlers.Login)
	API_GATEWAY.GET(path.LOGOUT, handlers.Logout)

	API_GATEWAY.PUT(path.AUTH_USER_UPGRADE, handlers.UpgradeUser)
	API_GATEWAY.PUT(path.AUTH_USER_DOWNGRADE, handlers.DowngradeUser)

	API_GATEWAY.Start(API_GATEWAY_PORT)
}
