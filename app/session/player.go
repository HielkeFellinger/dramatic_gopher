package session

import (
	"encoding/json"
	"log"
	"time"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/gorilla/websocket"
)

const (
	// Time units
	writeWait = 10 * time.Second
	readWait  = 60 * time.Second
	pongWait  = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type Player struct {
	Id          string
	DisplayName string
	conn        *websocket.Conn
	gameSession *GameSession
	send        chan []byte
}

func initPlayer(Id string, conn *websocket.Conn, GameSession *GameSession) *Player {
	return &Player{
		Id:          Id,
		conn:        conn,
		gameSession: GameSession,
		send:        make(chan []byte, 256),
	}
}

func (p *Player) readPump() {
	defer func() {
		_ = p.conn.Close()
	}()

	//_ = p.Conn.SetReadDeadline(time.Now().Add(readWait))
	//p.Conn.SetPongHandler(func(string) error {
	//	_ = p.Conn.SetReadDeadline(time.Now().Add(readWait))
	//	return nil
	//})

	// https://github.com/gorilla/websocket/blob/main/examples/chat/client.go

	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// Attempt to parse message
		var parsedMessage models.BasicRequestMessage
		log.Printf("received: %s", message)
		if err = json.Unmarshal(message, &parsedMessage); err != nil {
			log.Println("Could not parse message:", err)
			break
		}
		parsedMessage.UserId = p.Id

		// Check Channel & Send
		if p.gameSession != nil && p.gameSession.Events != nil {
			p.gameSession.Events <- parsedMessage
		}
	}
}

// WritePump sends messages from the `send` channel to the WebSocket connection
func (p *Player) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = p.conn.Close()
	}()
	for {
		select {
		case message, ok := <-p.send:
			// @TODO fix ignoring returns
			_ = p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := p.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = writer.Write(message)

			if CloseErr := writer.Close(); CloseErr != nil {
				return
			}
		case <-ticker.C:
			pingErr := p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if pingErr != nil {
				log.Printf("Error writing ping: %v", pingErr)
				return
			}
		}
	}
}
