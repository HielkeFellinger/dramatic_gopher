package ecs

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs/core"
	"github.com/google/uuid"
)

type Entity interface {
	GetId() uuid.UUID
	GetComponents() []*core.Component
}

type BaseEntity struct {
	Id         uuid.UUID         `json:"id"`
	Components []*core.Component `json:"core"`
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		Id:         uuid.New(),
		Components: make([]*core.Component, 0),
	}
}

func (b BaseEntity) GetId() uuid.UUID {
	return b.Id
}

func (b BaseEntity) GetComponents() []*core.Component {
	return b.Components
}
