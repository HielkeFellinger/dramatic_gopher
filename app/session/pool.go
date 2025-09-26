package session

import "github.com/HielkeFellinger/dramatic_gopher/app/engine"

var runningSessionPool = initGamePool()

type gamePool struct {
	Register   chan *GameSession
	Unregister chan *GameSession
	Sessions   map[*GameSession]bool
}

func initGamePool() *gamePool {
	gp := &gamePool{
		Register:   make(chan *GameSession),
		Unregister: make(chan *GameSession),
		Sessions:   make(map[*GameSession]bool),
	}
	go gp.Run()

	return gp
}

func (gp *gamePool) Run() {
	for {
		select {
		case session := <-gp.Register:
			gp.Sessions[session] = true
		case session := <-gp.Unregister:
			delete(gp.Sessions, session)
		}
	}
}

func (gp *gamePool) getRunningSessionByGameId(id string) *GameSession {
	for session, _ := range runningSessionPool.Sessions {
		game := *session.Game
		if game.GetId() == id {
			return session
		}
	}
	return nil
}

func AddGameToPool(leadId string, leadName string, game *engine.BaseGame) {
	// Create & Add Lead user
	newSession := initGameSession(leadId, leadName, game)
	go newSession.Run()

	// Register
	runningSessionPool.Register <- newSession
}

func IsGameRunning(id string) bool {
	for session, _ := range runningSessionPool.Sessions {
		game := *session.Game
		if game.GetId() == id {
			return true
		}
	}
	return false
}

func AddUserIdAndNameToAccessGame(userId string, userName string, gameId string) bool {
	for session, _ := range runningSessionPool.Sessions {
		if session.Game.GetId() == gameId {
			session.addUserIdAllowedToAccessSession(userId, userName)
			return true
		}
	}
	return false
}

func GetRunningGamePointer(id string) *engine.BaseGame {
	for session, _ := range runningSessionPool.Sessions {
		if session.Game.GetId() == id {
			return session.Game
		}
	}
	return nil
}

func IsUserIdAllowedToAccessGame(userId string, gameId string) bool {
	for session, _ := range runningSessionPool.Sessions {
		if session.Game.GetId() == gameId {
			return session.isUserIdAllowedToAccessSession(userId)
		}
	}
	return false
}

func IsUserIdLeadInGameById(userId string, gameId string) bool {
	for session, _ := range runningSessionPool.Sessions {
		game := *session.Game
		if game.GetId() == gameId {
			return session.isUserIdLeadInSession(userId)
		}
	}
	return false
}
