package initializers

import (
	"database/sql"
	"log"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/utils"
	_ "github.com/mattn/go-sqlite3"
)

func LoadDatabase() {
	log.Println("INIT: Attempting connecting to / creating Database")
	db, err := sql.Open("sqlite3", config.CurrentConfig.DatabaseFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Assure default tables are created
	loadDefaultTables(db)

	// Assure default content has been set
	loadDefaultContent(db)

	// Save connection
	models.DB = db
}

func loadDefaultTables(db *sql.DB) {
	log.Println("	Ensuring default tables are available")
	// Set default Tables
	sqlUsersStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        display_name TEXT NOT NULL,
        password TEXT NOT NULL
    );
    `
	if _, err := db.Exec(sqlUsersStmt); err != nil {
		log.Fatal(err)
	}
	log.Println("	Table 'users' is loaded successfully")

	sqlCampaignsStmt := `
    CREATE TABLE IF NOT EXISTS campaigns (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL UNIQUE 
    );
    `
	if _, err := db.Exec(sqlCampaignsStmt); err != nil {
		log.Fatal(err)
	}
	log.Println("	Table 'campaigns' is loaded successfully")

	sqlCampaignAccessStmt := `
    CREATE TABLE IF NOT EXISTS campaign_access (
        id INTEGER PRIMARY KEY,
        user_id INTEGER,
        campaign_id INTEGER,
        role TEXT NOT NULL,
	    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
	    FOREIGN KEY(campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
    `
	if _, err := db.Exec(sqlCampaignAccessStmt); err != nil {
		log.Fatal(err)
	}
	log.Println("	Table 'campaign_access' is loaded successfully")
}

func loadDefaultContent(db *sql.DB) {
	log.Println("	Ensuring default data is available")

	sqlEnsureDefaultAdminStmt := `
	INSERT INTO users(id, name, display_name, password) 
	SELECT 1, 'admin', 'admin', ?
	WHERE NOT EXISTS (SELECT 1 FROM users WHERE id = 1);
	`
	defaultAdminPassword, err := utils.HashPassword(config.CurrentConfig.DefaultAdminPassword)
	if err != nil {
		log.Fatal(err)
	}
	if _, insertErr := db.Exec(sqlEnsureDefaultAdminStmt, string(defaultAdminPassword)); insertErr != nil {
		log.Fatal(insertErr)
	}
	log.Println("	'admin' is loaded successfully")
}
