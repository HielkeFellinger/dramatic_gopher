package ecs

import "github.com/google/uuid"

type World interface {
	GetEntities() []*Entity
	GetSystems() []*System
	GetItemEntities() []Entity
	GetCharacterEntities() []Entity
	GetMapEntities() []Entity
	GetStorageEntities() []Entity
}

type BaseWorld struct {
	Entities []*Entity
	Systems  []*System

	UuidToEntity          map[uuid.UUID]Entity
	UuidToItemEntity      map[uuid.UUID]Entity
	UuidToCharacterEntity map[uuid.UUID]Entity
	UuidToMapEntity       map[uuid.UUID]Entity
	UuidToStorageEntity   map[uuid.UUID]Entity
}

func (w *BaseWorld) GetEntities() []*Entity {
	//TODO implement me
	panic("implement me")
}

func (w *BaseWorld) GetSystems() []*System {
	//TODO implement me
	panic("implement me")
}

func NewBaseWorld() *BaseWorld {
	return &BaseWorld{
		UuidToEntity:          make(map[uuid.UUID]Entity),
		UuidToItemEntity:      make(map[uuid.UUID]Entity),
		UuidToCharacterEntity: make(map[uuid.UUID]Entity),
		UuidToMapEntity:       make(map[uuid.UUID]Entity),
		UuidToStorageEntity:   make(map[uuid.UUID]Entity),
	}
}

func (w *BaseWorld) GetItemEntities() []Entity {
	return w.getEntityValuesOfMap(w.UuidToItemEntity)
}

func (w *BaseWorld) GetCharacterEntities() []Entity {
	return w.getEntityValuesOfMap(w.UuidToCharacterEntity)
}

func (w *BaseWorld) GetMapEntities() []Entity {
	return w.getEntityValuesOfMap(w.UuidToMapEntity)
}

func (w *BaseWorld) GetStorageEntities() []Entity {
	return w.getEntityValuesOfMap(w.UuidToStorageEntity)
}

func (w *BaseWorld) getEntityValuesOfMap(dict map[uuid.UUID]Entity) []Entity {
	values := make([]Entity, len(dict))
	index := 0
	for _, value := range dict {
		values[index] = value
		index++
	}
	return values
}

func (w *BaseWorld) getFilteredEntityValuesOfMap(dict map[uuid.UUID]Entity, filter func(entity Entity) bool) []Entity {
	values := make([]Entity, 0)
	for _, value := range dict {
		if filter(value) {
			values = append(values, value)
		}
	}
	return values
}
