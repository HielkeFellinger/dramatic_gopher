package routes

import (
	"github.com/gin-gonic/gin"
	"hielkefellinger.nl/dramatic_gopher/app/pages"
)

func HandlePageRoutes(router *gin.Engine) {

	// Pages
	router.GET("/", pages.Homepage())
}
