package pages

import (
	"net/http"
	"time"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/utils"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoadLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications := getNotifications(c)
		user := c.MustGet("user").(models.User)

		err := render(c, http.StatusOK, views.LoginPage(user, notifications))
		if err != nil {
			return
		}
	}
}

func HandleLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load Info
		user := c.MustGet("user").(models.User)
		notifications := getNotifications(c)

		// Validate Login
		var loggingRequest struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}
		bindErr := c.Bind(&loggingRequest)
		if bindErr != nil {
			noBindErr := models.NewNotification(models.Error, "Could not parse request, please retry")
			notifications = append(notifications, noBindErr)
		}

		if len(notifications) == 0 {
			// Attempt to Authenticate
			authUser, _ := models.UserService.GetUserByUsername(loggingRequest.Username)

			// Check the validity of the password hash; if no user matches, the password is an invalid hash or has a cost of 0
			if cost, err := bcrypt.Cost([]byte(authUser.Password)); err != nil || cost == 0 {
				time.Sleep(5 * time.Second) // Ensures minimal duration in auth attempt
				notifications = append(notifications, models.NewNotification(models.Error, "Invalid username or password"))
			} else {
				if errBcrypt := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(loggingRequest.Password)); errBcrypt != nil {
					notifications = append(notifications, models.NewNotification(models.Error, "Invalid username or password"))
				} else {
					var authCookieContent = utils.SessionCookieContent{
						ID:          authUser.Id,
						Role:        authUser.Role,
						DisplayName: authUser.DisplayName,
					}
					if errCookie := utils.SetSessionJWTCookie(authCookieContent, c); errCookie != nil {
						notifications = append(notifications, models.NewNotification(models.Error, "Failed to create Cookie"))
					} else {
						// OK
						c.Redirect(http.StatusFound, "/game/load/")
						return
					}
				}
			}
		}

		// FAILURE
		err := render(c, http.StatusBadRequest, views.LoginPage(user, notifications))
		if err != nil {
			return
		}
	}
}

func LoadRegisterPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications := getNotifications(c)
		user := c.MustGet("user").(models.User)

		err := render(c, http.StatusOK, views.RegisterPage(user, notifications))
		if err != nil {
			return
		}
	}
}
