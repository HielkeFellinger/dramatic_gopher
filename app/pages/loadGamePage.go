package pages

import (
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/engine"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
)

func LoadGamePage() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Load the games
		engine.FindAvailableGames()

		err := render(c, http.StatusOK, views.LoadGamePage())
		if err != nil {
			return
		}
	}
}
