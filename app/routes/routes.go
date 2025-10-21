package routes

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/middleware"
	"github.com/HielkeFellinger/dramatic_gopher/app/pages"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/gin-gonic/gin"
)

func HandlePageRoutes(router *gin.Engine) {

	router.GET("/", middleware.EnsureUserValuesIsSet, pages.Homepage())

	// Pages
	router.GET("/game/new", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/load", middleware.EnsureUserIsLoggedIn, pages.LoadGamePage())
	router.GET("/game/join/:game_id", middleware.EnsureUserIsLoggedIn, pages.LoadJoinGamePage())
	router.POST("/game/join/:game_id", middleware.EnsureUserIsLoggedIn, pages.HandleJoinGame())
	router.GET("/game/session/:game_id", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, pages.LoadGameSessionPage())

	// Users
	router.GET("/user/login", middleware.EnsureUserValuesIsSet, pages.LoadLoginPage())
	router.POST("/user/login", middleware.EnsureUserValuesIsSet, pages.HandleLoginPage())

	// Session
	router.GET("/session/:game_id/ws", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, session.HandleWebsocket)
}
