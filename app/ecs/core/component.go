package core

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Component interface {
	ComponentType() ComponentType
	SetId(id uuid.UUID)
	GetType() ComponentType
	AllowMultipleInEntity() bool
	ParseToRawComponent() (RawComponent, error)
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

func (c *BaseComponent) ComponentType() ComponentType {
	// TODO: Implement
	return UnknownComponentType
}

func (c *BaseComponent) ParseToRawComponent() (RawComponent, error) {
	return RawComponent{}, errors.New(fmt.Sprintf("ParseToRawComponent() not implemented. on type: '%d'", c.ComponentType()))
}
