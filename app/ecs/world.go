package ecs

type World interface {
	GetEntities() []*Entity
	GetSystems() []*System
}

type BaseWorld struct {
	Entities []*Entity
	Systems  []*System
}

func NewBaseWorld() *BaseWorld {
	return &BaseWorld{}
}
