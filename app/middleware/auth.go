package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/HielkeFellinger/dramatic_gopher/app/utils"
	"github.com/gin-gonic/gin"
)

const loginPageLocation = "/user/login"

func EnsureUserValuesIsSet(c *gin.Context) {
	c.Set("user", ensureSessionCookieAndGetUpToDateUser(c))
	c.Next()
}

func EnsureUserIsLoggedIn(c *gin.Context) {
	user := ensureSessionCookieAndGetUpToDateUser(c)
	notifications := getNotifications(c)
	c.Set("user", user)

	// Test validity of user session
	id, err := strconv.ParseInt(user.Id, 10, 64)
	if err != nil {
		unauthorized(c, notifications)
		return
	}
	// Check existence of user
	authUser, _ := models.UserService.GetUserById(id)
	if authUser.Id == "0" || authUser.Id != user.Id {
		unauthorized(c, notifications)
		return
	}

	c.Set("user", authUser)
	c.Next()
}

func EnsureUserValueIsSetAndAllowedToAccessGame(c *gin.Context) {
	user := ensureSessionCookieAndGetUpToDateUser(c)
	notifications := getNotifications(c)
	c.Set("user", user)

	gameId := c.Param("game_id")
	if gameId == "" || !session.IsUserIdAllowedToAccessGame(user.Id, gameId) {
		unauthorized(c, notifications)
		return
	}

	c.Next()
}

func ensureSessionCookieAndGetUpToDateUser(c *gin.Context) models.User {
	// Retrieve or recover SessionContent
	sessionContent, sessionError := utils.ParseSessionCookie(c)
	if sessionError != nil {
		sessionContent = utils.NewSessionCookieContent()
	}

	// Update cookie
	setCookieErr := utils.SetSessionJWTCookie(sessionContent, c)
	if setCookieErr != nil {
		log.Printf("Could not set/update session cookie: '%v'", setCookieErr.Error())
	}

	// Return data
	return models.User{
		Id:          sessionContent.ID,
		Role:        sessionContent.Role,
		DisplayName: sessionContent.DisplayName,
	}
}

func unauthorized(c *gin.Context, notifications []models.Notification) {
	c.Set("notification", append(notifications, models.NewNotification(models.Error, "401 - Unauthorized")))
	c.Redirect(http.StatusFound, loginPageLocation)
	c.AbortWithStatus(http.StatusUnauthorized)
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
