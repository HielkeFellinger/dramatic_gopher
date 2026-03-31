package pages

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/engine"
	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
)

func LoadGamePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Info
		notifications := getNotifications(c)

		// Load the baseGames
		baseGames := engine.FindAvailableGames()
		games := make([]engine.Game, len(baseGames))
		for index, game := range baseGames {
			// Check if game is registered
			if campaign, err := models.CampaignService.LoadCampaignOfDataDir(game.DataDir); err == nil {
				log.Println(campaign)
				game.Id = campaign.Id
				game.Running = session.IsGameRunning(game.Id)
			} else {
				log.Printf("Game save dir '%s' could not be loaded from the database", game.DataDir)
			}
			games[index] = game
		}

		if err := render(c, http.StatusOK, views.LoadGamePage(games, notifications)); err != nil {
			return
		}
	}
}

func LoadJoinGamePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Info
		notifications := getNotifications(c)
		gameId := c.Param("game_id")
		user := c.MustGet("user").(models.User)

		game, loadErr := retrieveGame(gameId)
		if loadErr != nil {
			notifications = append(notifications, models.NewNotification(models.Error, "404 - Game does not exist"))
			saveNotifications(c, notifications)
			c.Redirect(http.StatusFound, "/game/load")
			return
		}

		if err := render(c, http.StatusOK, views.JoinGamePage(game, user, notifications)); err != nil {
			return
		}
	}
}

func HandleJoinGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Info
		user := c.MustGet("user").(models.User)
		gameId := c.Param("game_id")
		notifications := getNotifications(c)

		game, loadErr := retrieveGame(gameId)
		if loadErr != nil {
			notifications = append(notifications, models.NewNotification(models.Error, "404 - Game does not exist"))
			saveNotifications(c, notifications)
			c.Redirect(http.StatusFound, "/game/load")
			return
		}

		// Validate Request
		var joinGameRequest struct {
			DisplayName string `form:"displayName"`
			Password    string `form:"password"`
		}
		bindErr := c.Bind(&joinGameRequest)
		if bindErr != nil {
			noBindErr := models.NewNotification(models.Error, "Could not parse request, please retry")
			notifications = append(notifications, noBindErr)
		}

		if len(notifications) == 0 {
			// Attempt to authenticate
			//if game.AuthenticateAsLead(joinGameRequest.Password) {
			//	if !game.IsRunning() {
			//		// Start game - OK
			//		session.AddGameToPool(user.Id, joinGameRequest.DisplayName, game)
			//		c.Redirect(http.StatusFound, "/game/session/"+game.Id)
			//		return
			//	} else {
			//		if !session.IsUserIdLeadInGameById(user.Id, game.GetId()) {
			//			notifications = append(notifications,
			//				models.NewNotification(models.Error, "Another user has already joined this game as Lead!"))
			//		} else {
			//			// Join running game - OK
			//			c.Redirect(http.StatusFound, "/game/session/"+game.Id)
			//			return
			//		}
			//	}
			//} else if game.AuthenticateAsClient(joinGameRequest.Password) {
			//	if !game.IsRunning() {
			//		notifications = append(notifications,
			//			models.NewNotification(models.Error, "Game is not running, not allowed to start game if not lead!"))
			//	} else {
			//		if session.IsUserIdLeadInGameById(user.Id, game.GetId()) {
			//			notifications = append(notifications,
			//				models.NewNotification(models.Error, "User is already linked to game as Lead; using client password is blocked"))
			//		} else {
			//			// Join running game - OK
			//			if !session.AddUserIdAndNameToAccessGame(user.Id, joinGameRequest.DisplayName, game.GetId()) {
			//				notifications = append(notifications,
			//					models.NewNotification(models.Error, "Failed to add you to the game allow list, please retry"))
			//			} else {
			//				c.Redirect(http.StatusFound, "/game/session/"+game.Id)
			//				return
			//			}
			//		}
			//	}
			//} else {
			//	notifications = append(notifications, models.NewNotification(models.Error, "No valid credentials"))
			//}
		}

		// FAILURE
		if renderErr := render(c, http.StatusOK, views.JoinGamePage(game, user, notifications)); renderErr != nil {
			return
		}
	}
}

func LoadGameSessionPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Info
		user := c.MustGet("user").(models.User)
		gameId := c.Param("game_id")
		notifications := getNotifications(c)

		game, loadErr := retrieveGame(gameId)
		if loadErr != nil {
			notifications = append(notifications, models.NewNotification(models.Error, "404 - Game does not exist"))
			saveNotifications(c, notifications)
			c.Redirect(http.StatusFound, "/game/load")
			return
		}

		if err := render(c, http.StatusOK, views.Session(user, game)); err != nil {
			return
		}
	}
}

func RegisterGameData() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.MustGet("user").(models.User)
		dataDir := c.Param("data_dir")
		notifications := getNotifications(c)

		// Test if a match already exists and is locked to an account
		if campaign, err := models.CampaignService.LoadCampaignOfDataDir(dataDir); err == nil {
			notifications = append(notifications, models.NewNotification(models.Warning,
				fmt.Sprintf("201 - Game Dir is already linked to campaign: '%s'.", campaign.Name)))
			saveNotifications(c, notifications)
			c.Redirect(http.StatusFound, "/game/load")
		}

		// Check if dir is available as a game
		baseGames := engine.FindAvailableGames()
		var match *engine.BaseGame
		for _, game := range baseGames {
			if game.DataDir == dataDir {
				match = game
			}
		}
		if match == nil {
			notifications = append(notifications, models.NewNotification(models.Error,
				"404 - Game Dir does not exist or is not a valid game!"))
			saveNotifications(c, notifications)
			c.Redirect(http.StatusFound, "/game/load")
			return
		}

		// Allow to be converted
		rawCampaign := models.Campaign{Name: match.Title, Description: match.Description}
		if err := render(c, http.StatusOK, views.RegisterGameDirAsCampaign(rawCampaign, dataDir, notifications)); err != nil {
			return
		}
	}
}

func retrieveGame(gameId string) (*engine.BaseGame, error) {
	// Attempt to load from session
	if session.IsGameRunning(gameId) {
		if game := session.GetRunningGamePointer(gameId); game != nil {
			return game, nil
		}
	}

	// Attempt to load form file
	game, loadErr := engine.LoadGameByDirectoryName(gameId)
	return game, loadErr
}
