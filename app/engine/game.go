package engine

type Game interface {
	GetName() string
	GetDescription() string
}

type BaseGame struct {
}

func NewBaseGame() *BaseGame {
	return &BaseGame{}
}
