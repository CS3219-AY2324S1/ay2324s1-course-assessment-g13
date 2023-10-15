package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/utils/path"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const API_GATEWAY_PORT = ":1234"

func main() {
	config.ConnectDb()
	log.Println("Starting development server...")

	API_GATEWAY := echo.New()

	corsMiddleware := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodPut, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})

	API_GATEWAY.Use(corsMiddleware, handlers.RequireAuthenticationMiddleWare)

	API_GATEWAY.GET(path.GITHUB_LOGIN, handlers.GithubLogin)
	API_GATEWAY.GET(path.LOGOUT, handlers.Logout)
	API_GATEWAY.GET(path.REFRESH, handlers.Refresh)

	API_GATEWAY.GET(path.AUTH_USER, handlers.GetUser)
	API_GATEWAY.DELETE(path.AUTH_USER, handlers.DeleteUser)

	API_GATEWAY.GET("/", handlers.RootHandler)
	API_GATEWAY.GET("/github", handlers.GithubLoginHandler)

	API_GATEWAY.GET(path.AUTH_USER_UPGRADE, handlers.UpgradeUser)
	API_GATEWAY.GET(path.AUTH_USER_DOWNGRADE, handlers.DowngradeUser)

	API_GATEWAY.Any(path.ALL_USER_SERVICE, handlers.HandleUserService)
	API_GATEWAY.Any(path.ALL_QUESTION_SERVICE, handlers.HandleQuestionService)
	API_GATEWAY.Any(path.ALL_MATCHING_SERVICE, handlers.HandleMatchingService)
	API_GATEWAY.Any(path.ALL_COLLAB_SERVICE, handlers.HandleCollabService)

	API_GATEWAY.Start(API_GATEWAY_PORT)
}
