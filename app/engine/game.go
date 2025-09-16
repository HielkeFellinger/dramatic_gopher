package engine

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs"
	"golang.org/x/crypto/bcrypt"
)

type Game interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetImageUrl() string
	IsRunning() bool
	Init() error
	Validate() (string, error)
	AuthenticateAsLead(password string) bool
	AuthenticateAsClient(password string) bool
}

type GameCrypto struct {
	LeadPassword   string `json:"leadPassword"`
	ClientPassword string `json:"clientPassword"`
	Description    string `json:"description"`
}

type BaseGame struct {
	Id          string
	Title       string     `json:"title"`
	Version     string     `json:"-"` // Sourced form SafeFile
	Crypto      GameCrypto `json:"crypto"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"imageUrl"`
	SaveFile    string     `json:"saveFile"`
	Running     bool
	World       *ecs.World // Sourced form SafeFile
}

func (bg *BaseGame) Init() error {
	validSaveFile, validationErr := bg.Validate()
	if validationErr != nil {
		return fmt.Errorf("failed to validate campaign: '%q'", bg.Id)
	}

	// Attempt to initialize...
	log.Printf("Validating game file: '%s'", validSaveFile)

	return nil
}

func (bg *BaseGame) Validate() (string, error) {
	// Test if there are files to load in the first place
	if len(bg.SaveFile) == 0 {
		return "", fmt.Errorf("no game files found for campaign dir/id: '%q'", bg.Id)
	}

	campaignPath := filepath.Join(config.CurrentConfig.CampaignSavesDir, bg.Id)

	// Test if all game files are readable
	if validGameFilePath, gameFilePathErr := getValidatedGameFile(campaignPath, strings.TrimSpace(bg.SaveFile)); gameFilePathErr != nil {
		return "", gameFilePathErr
	} else {
		return validGameFilePath, nil
	}
}

func (bg *BaseGame) GetId() string {
	return bg.Id
}

func (bg *BaseGame) GetTitle() string {
	return bg.Title
}

func (bg *BaseGame) GetDescription() string {
	return bg.Description
}

func (bg *BaseGame) GetImageUrl() string {
	return bg.ImageUrl
}

func (bg *BaseGame) IsRunning() bool {
	return bg.Running
}

func (bg *BaseGame) AuthenticateAsClient(password string) bool {
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(bg.Crypto.ClientPassword), []byte(password))
	return errBcrypt == nil
}

func (bg *BaseGame) AuthenticateAsLead(password string) bool {
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(bg.Crypto.LeadPassword), []byte(password))
	return errBcrypt == nil
}
