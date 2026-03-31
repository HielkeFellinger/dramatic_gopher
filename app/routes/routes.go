package routes

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/HielkeFellinger/dramatic_gopher/app/middleware"
	"github.com/HielkeFellinger/dramatic_gopher/app/pages"
	"github.com/HielkeFellinger/dramatic_gopher/app/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func HandleNotificationSessionStore(router *gin.Engine) {
	cookieStore := cookie.NewStore([]byte(config.CurrentConfig.SessionSecret))
	router.Use(sessions.Sessions("dramatic_session", cookieStore))
}

func HandlePageRoutes(router *gin.Engine) {

	router.GET("/", middleware.EnsureUserValuesIsSet, pages.Homepage())

	// Pages
	router.GET("/game/new", middleware.EnsureUserValuesIsSet, pages.Homepage())
	router.GET("/game/load", middleware.EnsureUserIsLoggedIn, pages.LoadGamePage())
	router.GET("/game/join/:game_id", middleware.EnsureUserIsLoggedIn, pages.LoadJoinGamePage())
	router.POST("/game/join/:game_id", middleware.EnsureUserIsLoggedIn, pages.HandleJoinGame())
	router.GET("/game/session/:game_id", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, pages.LoadGameSessionPage())
	router.POST("/game/register/:data_dir", middleware.EnsureUserHasAdminRole, pages.HandleJoinGame())
	router.GET("/game/register/:data_dir", middleware.EnsureUserHasAdminRole, pages.RegisterGameData())

	// Users
	router.GET("/user/login", middleware.EnsureUserValuesIsSet, pages.LoadLoginPage())
	router.POST("/user/login", middleware.EnsureUserValuesIsSet, pages.HandleLoginPage())
	router.GET("/user/logout", middleware.EnsureUserValuesIsSet, pages.HandleLogoutPage())
	router.GET("/user/register", middleware.EnsureUserValuesIsSet, pages.LoadRegisterPage())
	router.POST("/user/register", middleware.EnsureUserValuesIsSet, pages.HandleRegisterPage())

	// Session
	router.GET("/session/:game_id/ws", middleware.EnsureUserValueIsSetAndAllowedToAccessGame, session.HandleWebsocket)
}
