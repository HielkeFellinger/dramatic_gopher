package core

type IdentityComponent struct {
	BaseComponent
	Name        string `json:"name"`
	Description string `json:"description"`
}
