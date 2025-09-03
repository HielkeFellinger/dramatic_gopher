package engine

import "encoding/json"

type RawSaveFile struct {
	Version    string      `json:"version"`
	Crypto     GameCrypto  `json:"crypto"`
	Items      []RawEntity `json:"items"`
	Characters []RawEntity `json:"characters"`
	Maps       []RawEntity `json:"maps"`
	Storage    []RawEntity `json:"storage"`
}

type RawEntity struct {
	Id         string         `json:"id"`
	Components []RawComponent `json:"components"`
}

type RawComponent struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}
