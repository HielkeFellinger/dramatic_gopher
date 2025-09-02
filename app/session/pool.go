package session

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

func IsGameRunning(id string) bool {
	for session, _ := range runningSessionPool.Sessions {
		game := *session.Game
		if game.GetId() == id {
			return true
		}
	}
	return false
}
