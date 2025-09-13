package pages

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func getNotifications(c *gin.Context) []models.Notification {
	var notifications = make([]models.Notification, 0)
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
	return notifications
}
