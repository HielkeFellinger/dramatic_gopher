package session

import (
	"log"
	"time"

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
	conn        *websocket.Conn
	GameSession *GameSession
}

func (p *Player) readPump() {
	defer func() {
		_ = p.conn.Close()
	}()

	p.conn.SetReadLimit(512)
	_ = p.conn.SetReadDeadline(time.Now().Add(readWait))
	p.conn.SetPongHandler(func(string) error {
		_ = p.conn.SetReadDeadline(time.Now().Add(readWait))
		return nil
	})

	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("received: %s", message)
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
		case <-ticker.C:
			pingErr := p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if pingErr != nil {
				log.Printf("Error writing ping: %v", pingErr)
				return
			}
		}
	}
}
