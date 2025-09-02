package engine

import "encoding/json"

// https://stackoverflow.com/questions/36525602/can-i-unmarshal-json-into-implementers-of-an-interface

type RawSaveFile struct {
	Version    string      `json:"version"`
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
