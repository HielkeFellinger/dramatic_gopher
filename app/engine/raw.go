package engine

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs/core"
)

type RawSaveFile struct {
	Version    string           `json:"version"`
	DateTime   string           `json:"datetime"`
	Items      []core.RawEntity `json:"items"`
	Characters []core.RawEntity `json:"characters"`
	Maps       []core.RawEntity `json:"maps"`
	Storage    []core.RawEntity `json:"storage"`
}
