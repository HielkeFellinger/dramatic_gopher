package ecs

import "github.com/HielkeFellinger/dramatic_gopher/app/ecs/components"

type Component interface {
	GetType() components.ComponentType
	AllowMultipleInEntity() bool
}
