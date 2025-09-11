package models

import "github.com/google/uuid"

type NotificationType string

const (
	Error   NotificationType = "ERROR"
	Warning NotificationType = "WARNING"
	Success NotificationType = "SUCCESS"
)

type Notification struct {
	Id      string
	Type    NotificationType
	Content string
}

func NewNotification(NotificationType NotificationType, content string) Notification {
	return Notification{
		Id:      uuid.NewString(),
		Type:    NotificationType,
		Content: content,
	}
}
