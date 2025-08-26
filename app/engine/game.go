package engine

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs"
)

type Game interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetImageUrl() string
	Init() error
	Validate() error
}

type BaseGame struct {
	Id          string
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	ImageUrl    string   `yaml:"image_url"`
	GameFiles   []string `yaml:"game_files"`
	World       *ecs.World
}

func NewBaseGame(id string) *BaseGame {
	return &BaseGame{
		Id:        id,
		GameFiles: make([]string, 0),
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

func (bg *BaseGame) Init() error {
	valid, validGameFiles, validationErrs := bg.Validate()
	if !valid || len(validationErrs) > 0 {
		return fmt.Errorf("failed to validate campaign: '%q'", bg.Id)
	}

	// Attempt to initialize...
	for _, validGameFile := range validGameFiles {
		// TODO LOAD!
		log.Printf("Validating game file: '%s'", validGameFile)
	}

	return nil
}

func (bg *BaseGame) Validate() (bool, []string, []error) {
	validGameFiles := make([]string, 0)
	valErrors := make([]error, 0)

	// Test if there are files to load in the first place
	if len(bg.GameFiles) == 0 {
		valErrors = append(valErrors, fmt.Errorf("no game files found for campaign dir/id: '%q'", bg.Id))
	}

	campaignPath := filepath.Join(config.CurrentConfig.CampaignSavesDir, bg.Id)

	// Test if all game files are readable
	for _, gameFilePath := range bg.GameFiles {
		if validGameFilePath, gameFilePathErr := getValidatedGameFile(campaignPath, strings.TrimSpace(gameFilePath)); gameFilePathErr != nil {
			valErrors = append(valErrors, gameFilePathErr)
		} else {
			validGameFiles = append(validGameFiles, validGameFilePath)
		}
	}

	return len(valErrors) == 0, validGameFiles, valErrors
}
