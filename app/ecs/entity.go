package ecs

import (
	"github.com/google/uuid"
)

type Entity interface {
	GetId() uuid.UUID
}

type BaseEntity struct {
	Id         uuid.UUID    `json:"id"`
	Components []*Component `json:"components"`
}

func NewBaseEntity() *BaseEntity {

	return &BaseEntity{
		Id: uuid.New(),
	}
}
