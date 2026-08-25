package core

import "encoding/json"

type RawEntity struct {
	Id         string         `json:"id"`
	Components []RawComponent `json:"core"`
}

type RawComponent struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}
