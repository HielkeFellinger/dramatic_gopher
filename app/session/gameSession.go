package session

import "github.com/HielkeFellinger/dramatic_gopher/app/engine"

type GameSession struct {
	Game       *engine.Game
	Register   chan *Player
	Unregister chan *Player
	Players    map[*Player]bool
}

func initGameSession() *GameSession {
	return &GameSession{
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Players:    make(map[*Player]bool),
	}
}

func (gs *GameSession) Run() {
	for {
		select {
		case player := <-gs.Register:
			gs.Players[player] = true
		case player := <-gs.Unregister:
			delete(gs.Players, player)

		}
	}
}
