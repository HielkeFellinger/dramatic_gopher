package engine

import "github.com/HielkeFellinger/dramatic_gopher/app/ecs"

type Game interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetImageUrl() string
}

type BaseGame struct {
	Id          string
	Title       string
	Description string
	ImageUrl    string
	GameFile    string
	World       *ecs.World
}

func NewBaseGame() *BaseGame {
	return &BaseGame{}
}

func (bg *BaseGame) GetId() string {
	return bg.Id
}

func (bg *BaseGame) GetTitle() string {
	return bg.Title
}

func (bg *BaseGame) GetDescription() string {
	return bg.Description
}

func (bg *BaseGame) GetImageUrl() string {
	return bg.ImageUrl
}
