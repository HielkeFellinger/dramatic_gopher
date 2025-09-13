package session

import (
	"log"
	"time"

	"github.com/HielkeFellinger/dramatic_gopher/app/engine"
)

type GameSession struct {
	Game       *engine.BaseGame
	Register   chan *Player
	Unregister chan *Player
	Events     chan requestMessage
	Players    map[*Player]bool
	Lead       *Player

	leadId               string
	authenticatedUserIds map[string]string
	playerIdToPlayer     map[string]*Player
}

func initGameSession(leadId string, leadName string, game *engine.BaseGame) *GameSession {
	session := GameSession{
		Game:                 game,
		Register:             make(chan *Player),
		Unregister:           make(chan *Player),
		Events:               make(chan requestMessage),
		Players:              make(map[*Player]bool),
		leadId:               leadId,
		authenticatedUserIds: make(map[string]string),
		playerIdToPlayer:     make(map[string]*Player),
	}
	session.authenticatedUserIds[leadId] = leadName
	return &session
}

func (gs *GameSession) Run() {
	checkIfEmptyTimer := time.NewTicker(180 * time.Second)
	defer checkIfEmptyTimer.Stop()
	for {
		select {
		case player := <-gs.Register:
			// Add Player
			gs.Players[player] = true
			gs.playerIdToPlayer[player.Id] = player

			// Add possible root
			if gs.leadId == player.Id {
				gs.Lead = player
			}
		case player := <-gs.Unregister:
			delete(gs.Players, player)
			delete(gs.playerIdToPlayer, player.Id)
			// @TODO Check removal lead player

		case message := <-gs.Events:
			// @TODO Handle Events
			log.Println(message)

		case <-checkIfEmptyTimer.C:
			// Clear session if session is empty during timeout check
			if len(gs.Players) == 0 {
				runningSessionPool.Unregister <- gs
				return
			}
		}
	}
}

func (gs *GameSession) isUserIdAllowedToAccessSession(id string) bool {
	_, ok := gs.authenticatedUserIds[id]
	return ok
}

func (gs *GameSession) isUserIdLeadInSession(id string) bool {
	return id == gs.leadId
}

func (gs *GameSession) addUserIdAllowedToAccessSession(id string, name string) {
	gs.authenticatedUserIds[id] = name
}
