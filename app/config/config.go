package config

import (
	"log"
	"os"
)

var CurrentConfig *Config

type Config struct {
	CryptCost            string
	JwtSecret            string
	Host                 string
	Port                 string
	CampaignSavesDir     string
	DatabaseFilePath     string
	DefaultAdminPassword string
	DefaultGamePassword  string
}

func InitConfig() *Config {
	if CurrentConfig != nil {
		return CurrentConfig
	}

	log.Println("INIT: Attempting Load new Config")
	configInit := &Config{}

	// Required
	configInit.Port = os.Getenv("PORT")
	configInit.JwtSecret = os.Getenv("JWT_SECRET")
	configInit.CampaignSavesDir = os.Getenv("CAMPAIGN_SAVES_DIR")
	configInit.DatabaseFilePath = os.Getenv("DATABASE_FILE_PATH")
	configInit.DefaultAdminPassword = os.Getenv("DEFAULT_ADMIN_PASSWORD")
	configInit.DefaultGamePassword = os.Getenv("DEFAULT_GAME_PASSWORD")

	if len(configInit.DefaultAdminPassword) < 5 {
		log.Panic("CONFIG: Invalid (to small <5 chars) or empty Default Admin Password (env.DEFAULT_ADMIN_PASSWORD)")
	}
	if len(configInit.DefaultGamePassword) < 5 {
		log.Panic("CONFIG: Invalid (to small <5 chars) or empty Default Game Password (env.DEFAULT_GAME_PASSWORD)")
	}
	if len(configInit.JwtSecret) < 8 {
		log.Panic("CONFIG: Invalid (to small <8 chars) or empty JWT secret (env.JWT_SECRET)")
	}
	if len(configInit.Port) == 0 {
		log.Panic("CONFIG: Invalid or missing Port (env.PORT)")
	}

	_, readError := os.ReadDir(configInit.CampaignSavesDir)
	if readError != nil {
		log.Panic("CONFIG: Error reading save dir (env.CAMPAIGN_SAVES_DIR): ", readError)
	}

	// Optional
	configInit.Host = os.Getenv("HOST")            // Allowed Empty
	configInit.CryptCost = os.Getenv("CRYPT_COST") // Will fall back to a set minimum

	CurrentConfig = configInit
	return CurrentConfig
}
