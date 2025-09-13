package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	//jsonReturn := gin.H{}

	// Checks

	// Upgrade Connection
	//conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	jsonReturn["error"] = err.Error()
	//	c.JSON(http.StatusBadRequest, jsonReturn)
	//	return
	//}

	// Attempt to load game (if not already started

	// Start Player pumps

}
