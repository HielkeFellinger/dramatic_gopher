package session

import "github.com/gin-gonic/gin"

type requestMessage struct {
	UserId  string           `json:"-"`
	Context *gin.Context     `json:"-"`
	Type    string           `json:"type"`
	Value   string           `json:"value"`
	Headers requestHxHeaders `json:"HEADERS"`
}

type requestHxHeaders struct {
	HxRequest     string `json:"HX-Request"`
	HxTrigger     string `json:"HX-Trigger"`
	HxTriggerName string `json:"HX-Trigger-Name"`
	HxTarget      string `json:"HX-Target"`
	HxCurrentUrl  string `json:"HX-Current-URL"`
}
