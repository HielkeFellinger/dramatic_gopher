package utils

import (
	"errors"
	"fmt"
	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const SessionCookieName = "Session"

type SessionCookieContent struct {
	ID        string
	ExpiresAt int64
}

func NewSessionCookieContent() SessionCookieContent {
	return SessionCookieContent{
		ID:        uuid.NewString(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
}

func SetSessionJWTCookie(content SessionCookieContent, c *gin.Context) error {
	// Override / Set default expiration
	content.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	// Generate a JWT token and Sign
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ID":        content.ID,
		"ExpiresAt": content.ExpiresAt,
	})
	tokenString, err := token.SignedString([]byte(config.CurrentConfig.JwtSecret))

	if err == nil {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(SessionCookieName, tokenString, 3600*24, "", "", false, false)
	}

	return err
}

func ParseSessionCookie(c *gin.Context) (SessionCookieContent, error) {
	var session SessionCookieContent

	tokenString, err := c.Cookie("Session")
	if err == nil {
		// Parse tokenString
		token, jwtErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what is expected
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.CurrentConfig.JwtSecret), nil
		})

		// Send jwtErr if failure in parsing
		if jwtErr != nil {
			return session, jwtErr
		}

		// Validate the cookie claims, expiration, and session content
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the expiration.
			if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64) {
				return session, errors.New("cookie expired; needs refresh")
			}

			// Get session, if exists
			session = SessionCookieContent{
				ID:        claims["ID"].(string),
				ExpiresAt: int64(claims["ExpiresAt"].(float64)),
			}
		}
	}
	return session, err
}

func ResetCookie(name string, c *gin.Context) {
	c.SetCookie(name, "", -1, "", "", false, false)
}
