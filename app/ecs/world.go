package ecs

type World interface {
	GetEntities() []*Entity
	GetSystems() []*System
}

type BaseWorld struct {
}

func NewBaseWorld() *BaseWorld {
	return &BaseWorld{}
}
