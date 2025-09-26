package session

import (
	"log"
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,

	// @TODO: SEC FAIL/DANGER THIS DOES BYPASS ORIGIN CHECK!!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebsocket(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	gameId := c.Param("game_id")
	jsonReturn := gin.H{}

	log.Printf("Attempting to connect to connect Player: '%s' to Game '%s'", user.Id, gameId)

	// Checks Access and if game is running
	if !IsUserIdAllowedToAccessGame(user.Id, gameId) {
		jsonReturn["error"] = "401 - Unauthorized"
		c.JSON(http.StatusUnauthorized, jsonReturn)
		return
	}

	// Get the session
	runningSession := runningSessionPool.getRunningSessionByGameId(gameId)
	if runningSession == nil {
		jsonReturn["error"] = "404 - Session not available"
		c.JSON(http.StatusNotFound, jsonReturn)
		return
	}

	// Upgrade Connection
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		jsonReturn["error"] = err.Error()
		c.JSON(http.StatusBadRequest, jsonReturn)
		return
	}
	ws.SetReadLimit(maxMessageSize)

	// Create new player
	player := &Player{
		Id:          user.Id,
		conn:        ws,
		gameSession: runningSession,
	}

	// Start Read and Write Pumps & register player
	log.Printf("Attempting to start read and write pumps of Player: '%s' to Game '%s'", user.Id, gameId)
	go player.readPump()
	go player.writePump()
	runningSession.Register <- player
}
