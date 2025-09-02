package pages

import (
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/engine"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
)

func LoadGamePage() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Load the baseGames
		baseGames := engine.FindAvailableGames()
		games := make([]engine.Game, len(baseGames))
		for index, game := range baseGames {
			game.Running = session.IsGameRunning(game.Id)
			games[index] = game
		}

		err := render(c, http.StatusOK, views.LoadGamePage(games))
		if err != nil {
			return
		}
	}
}
