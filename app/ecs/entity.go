package ecs

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs/components"
	"github.com/google/uuid"
)

type Entity interface {
	GetId() uuid.UUID
}

type BaseEntity struct {
	Id         uuid.UUID               `json:"id"`
	Components []*components.Component `json:"components"`
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		Id: uuid.New(),
	}
}
