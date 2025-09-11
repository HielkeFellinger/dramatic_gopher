package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
)

func FindAvailableGames() []*BaseGame {
	possibleGames := make([]*BaseGame, 0)
	campaignSaveDir := config.CurrentConfig.CampaignSavesDir

	entries, readErr := os.ReadDir(campaignSaveDir)
	if readErr != nil {
		log.Println("Could not read campaign saves dir: ", readErr)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if possibleGame, loadErr := LoadGameById(entry.Name()); loadErr == nil {
				possibleGames = append(possibleGames, possibleGame)
			} else {
				log.Println("Could not read potential campaign dir: ", loadErr)
			}
		}
	}

	return possibleGames
}

func LoadGameById(id string) (*BaseGame, error) {
	campaignSaveDir := config.CurrentConfig.CampaignSavesDir
	currentCampaignPath := filepath.Join(campaignSaveDir, id)
	if campaignFiles, campaignReadError := os.ReadDir(currentCampaignPath); campaignReadError == nil {
		for _, campaignEntry := range campaignFiles {
			if campaignEntry.IsDir() {
				continue
			}

			// A file has been found; attempt to load it!
			if campaignEntry.Name() == "game_info.json" {
				// Attempt to load basics:
				gameInfoData, gameReadErr := os.ReadFile(filepath.Join(currentCampaignPath, campaignEntry.Name()))
				if gameReadErr != nil {
					return nil, fmt.Errorf("error Reading campaign 'game_info.json' File: '%s'", gameReadErr.Error())
				}

				var potentialGame BaseGame
				log.Println(string(gameInfoData))
				if gameUnmarshallErr := json.Unmarshal(gameInfoData, &potentialGame); gameUnmarshallErr != nil {
					return nil, fmt.Errorf("error parsing campaign 'game_info.json' File: '%s'", gameUnmarshallErr.Error())
				}

				// Validate if all files are reachable!
				potentialGame.Id = id
				return &potentialGame, nil
			}
		}
	} else {
		return nil, fmt.Errorf("could not read campaign saves dir: %s", campaignReadError.Error())
	}
	return nil, fmt.Errorf("could not load campaign by id: '%s', no match found", id)
}
