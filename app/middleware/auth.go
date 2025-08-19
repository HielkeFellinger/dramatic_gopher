package middleware

import (
	"log"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/utils"
	"github.com/gin-gonic/gin"
)

func EnsureUserValuesIsSet(c *gin.Context) {
	c.Set("user", ensureSessionCookieAndGetUpToDateUser(c))
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
		Id: sessionContent.ID,
	}
}
