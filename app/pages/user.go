package pages

import (
	"log"
	"net/http"
	"strings"
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
		if err := render(c, http.StatusOK, views.LoginPage(user, notifications)); err != nil {
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
						c.Redirect(http.StatusFound, "/")
						return
					}
				}
			}
		}

		// FAILURE
		if err := render(c, http.StatusBadRequest, views.LoginPage(user, notifications)); err != nil {
			return
		}
	}
}

func HandleLogoutPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := "unknown"
		if rawUser, exists := c.Get("user"); exists {
			user := rawUser.(models.User)
			username = user.Name
		}

		// Attempt Reset Cookie
		utils.ResetCookie(utils.SessionCookieName, c)
		log.Printf("User: '%s' has been logged out", username)

		c.Redirect(http.StatusFound, "/")
		return
	}
}

func LoadRegisterPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications := getNotifications(c)
		user := c.MustGet("user").(models.User)
		if err := render(c, http.StatusOK, views.RegisterPage(user, notifications)); err != nil {
			return
		}
	}
}

func HandleRegisterPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications := getNotifications(c)

		// Validate Registration
		var registerRequest struct {
			Username      string `form:"username"`
			Password      string `form:"password"`
			DisplayName   string `form:"displayName"`
			PasswordCheck string `form:"passwordCheck"`
		}
		if bindErr := c.Bind(&registerRequest); bindErr != nil {
			notifications = append(notifications, models.NewNotification(models.Error, "Could not parse registration request, please retry"))
		}

		registerRequest.Username = strings.TrimSpace(registerRequest.Username)
		registerRequest.DisplayName = strings.TrimSpace(registerRequest.DisplayName)

		user := models.User{
			Name:        registerRequest.Username,
			DisplayName: registerRequest.DisplayName,
			Password:    registerRequest.Password,
		}

		// Check Input Validity
		if !utils.CheckIfNameMeetsPolicy(registerRequest.Username) {
			notifications = append(notifications, models.NewNotification(models.Error, "Username does not meet policy"))
		}
		if !utils.CheckIfNameMeetsPolicy(registerRequest.DisplayName) {
			notifications = append(notifications, models.NewNotification(models.Error, "DisplayName does not meet policy"))
		}
		if registerRequest.Password != registerRequest.PasswordCheck {
			notifications = append(notifications, models.NewNotification(models.Error, "Password does not match"))
		}
		if !utils.CheckIfPasswordMeetsPolicy(registerRequest.Password, registerRequest.Username) {
			notifications = append(notifications, models.NewNotification(models.Error, "Password does not meet policy"))
		}

		// Check if user exists
		userMatch, _ := models.UserService.GetUserByUsername(registerRequest.Username)
		if userMatch.Name == registerRequest.Username {
			notifications = append(notifications, models.NewNotification(models.Error, "Username is unavailable"))
		}

		// OK so far
		if len(notifications) == 0 {
			log.Printf("Attempting to insert User: '%s'", registerRequest.Username)
			insertedUser, err := models.UserService.InsertUser(user)
			if err != nil {
				log.Printf("Failed to insert User: '%s'. Err: '%s'", registerRequest.Username, err.Error())
				notifications = append(notifications, models.NewNotification(models.Error, "Could not insert user, please try again"))
			} else {
				// OK
				log.Printf("User: '%s', ID: '%s' has been registered", insertedUser.Name, insertedUser.Id)
				c.Redirect(http.StatusFound, "/user/login")
				return
			}
		}

		// FAILURE
		if err := render(c, http.StatusBadRequest, views.RegisterPage(user, notifications)); err != nil {
			return
		}
	}
}
