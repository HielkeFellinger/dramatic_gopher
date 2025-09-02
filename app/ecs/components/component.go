package components

import "github.com/google/uuid"

type Component interface {
	SetId(id uuid.UUID)
	GetType() ComponentType
	AllowMultipleInEntity() bool
}

type BaseComponent struct {
	Id   uuid.UUID     `json:"-"`
	Type ComponentType `json:"-"`
}

func (c *BaseComponent) SetId(id uuid.UUID) {
	c.Id = id
}

func (c *BaseComponent) GetType() ComponentType {
	return c.Type
}

func (c *BaseComponent) AllowMultipleInEntity() bool {
	return false
}
