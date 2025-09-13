package routes

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/middleware"
	"github.com/HielkeFellinger/dramatic_gopher/app/pages"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/gin-gonic/gin"
)

func HandlePageRoutes(router *gin.Engine) {

	// Pages
	router.GET("/", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/new", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/load", middleware.EnsureUserValuesIsSet, pages.LoadGamePage())
	router.GET("/game/join/:game_id", middleware.EnsureUserValuesIsSet, pages.LoadJoinGamePage())
	router.POST("/game/join/:game_id", middleware.EnsureUserValuesIsSet, pages.HandleJoinGame())
	router.GET("/game/session/:game_id", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, pages.LoadGameSessionPage())

	// Session
	router.GET("/session/:game_id/ws", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, session.HandleWebsocket)
}
