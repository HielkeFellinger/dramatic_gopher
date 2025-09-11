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

func LoadJoinGamePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Game
		id := c.Param("id")
		game, loadErr := engine.LoadGameById(id)

		if loadErr != nil {
			_ = c.Error(loadErr)
			c.Abort()
		} else {
			game.Running = session.IsGameRunning(game.Id)
		}

		err := render(c, http.StatusOK, views.JoinGamePage(game))
		if err != nil {
			return
		}
	}
}
