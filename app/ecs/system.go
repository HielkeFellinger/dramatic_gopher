package ecs

type System interface {
}

type BaseSystem struct {
}

func NewBaseSystem() *BaseSystem {
	return &BaseSystem{}
}
