package pages

import (
	"encoding/json"
	"log"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func getNotifications(c *gin.Context) []models.Notification {
	var notifications = make([]models.Notification, 0)

	// Check context
	rawNotification, hasNotification := c.Get("notification")
	queryNotification := c.Query("notification")
	if hasNotification {
		notification := rawNotification.(models.Notification)
		notifications = append(notifications, notification)
	} else {
		if len(queryNotification) > 0 {
			notification := models.NewNotification(models.Error, queryNotification)
			notifications = append(notifications, notification)
		}
	}

	// Check Session flashes for notifications
	session := sessions.Default(c)
	sessionNotifications := session.Get("notifications")
	session.Set("notifications", "[]")
	_ = session.Save()
	var sessionNotifs []models.Notification
	log.Printf("-----> %v", sessionNotifications)
	if sessionNotifications != nil {
		if err := json.Unmarshal([]byte(sessionNotifications.(string)), &sessionNotifs); err == nil {
			notifications = append(notifications, sessionNotifs...)
		}
	}

	return notifications
}

func saveNotifications(c *gin.Context, notifications []models.Notification) {
	var session = sessions.Default(c)

	// Save to JSON
	if data, err := json.Marshal(notifications); err == nil {
		session.Set("notifications", string(data))
		_ = session.Save()
	}
}
