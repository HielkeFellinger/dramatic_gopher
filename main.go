package main

import (
	"fmt"
	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/HielkeFellinger/dramatic_gopher/app/initializers"
	"github.com/HielkeFellinger/dramatic_gopher/app/routes"
	"github.com/gin-gonic/gin"
	"log"
)

var engine *gin.Engine

func init() {
	log.Println("INIT: Starting Initialisation of Dramatic Gopher...")
	initializers.LoadEnvVariables()
	config.InitConfig()
	log.Println("INIT: Done. Initialisation Finished...")
}

func main() {
	loadGinEngine()

	// Serve Content
	log.Println("MAIN: Starting Gin.Engine")
	log.Fatal(engine.Run(fmt.Sprintf("%s:%s", config.CurrentConfig.Host, config.CurrentConfig.Port)))
}

func loadGinEngine() {
	log.Println("MAIN: Creation of Gin.Engine")
	engine = gin.Default()
	// Load Routes and (static) content
	log.Println("MAIN: Loading (Static) Content, Templates and Routes")
	routes.HandlePageRoutes(engine)
}
