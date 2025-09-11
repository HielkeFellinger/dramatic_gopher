package routes

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/middleware"
	"github.com/HielkeFellinger/dramatic_gopher/app/pages"
	"github.com/gin-gonic/gin"
)

func HandlePageRoutes(router *gin.Engine) {

	// Pages
	router.GET("/", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/new", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/load", middleware.EnsureUserValuesIsSet, pages.LoadGamePage())
	router.GET("/game/join/:id", middleware.EnsureUserValuesIsSet, pages.LoadJoinGamePage())
	router.POST("/game/join/:id", middleware.EnsureUserValuesIsSet, pages.LoadGamePage())
	router.GET("/session/:id", middleware.EnsureUserValuesIsSet, pages.LoadGamePage())
}
