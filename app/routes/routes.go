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
}
