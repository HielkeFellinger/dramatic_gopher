package models

import "github.com/gin-gonic/gin"

type BasicRequestMessage struct {
	UserId     string           `json:"-"`
	Context    *gin.Context     `json:"-"`
	RawMessage string           `json:"-"`
	Type       string           `json:"type"`
	Value      string           `json:"value"`
	Headers    RequestHxHeaders `json:"HEADERS"`
}

type RequestHxHeaders struct {
	HxRequest     string `json:"HX-Request"`
	HxTrigger     string `json:"HX-Trigger"`
	HxTriggerName string `json:"HX-Trigger-Name"`
	HxTarget      string `json:"HX-Target"`
	HxCurrentUrl  string `json:"HX-Current-URL"`
}
