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

	log.Printf("Reading Campaign data dir: '%s'", campaignSaveDir)
	entries, readErr := os.ReadDir(campaignSaveDir)
	if readErr != nil {
		log.Println("Could not read campaign data dir: ", readErr)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if possibleGame, loadErr := LoadGameByDirectoryName(entry.Name()); loadErr == nil {
				possibleGames = append(possibleGames, possibleGame)
			} else {
				log.Printf("Could not read potential campaign save dir: '%s'", loadErr)
			}
		}
	}
	return possibleGames
}

func LoadGameByDirectoryName(dirName string) (*BaseGame, error) {
	campaignSaveDir := config.CurrentConfig.CampaignSavesDir
	currentCampaignPath := filepath.Join(campaignSaveDir, dirName)

	log.Printf("Attempting to read Campaign save dir: '%s'", currentCampaignPath)
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
				if gameUnmarshallErr := json.Unmarshal(gameInfoData, &potentialGame); gameUnmarshallErr != nil {
					return nil, fmt.Errorf("error parsing campaign 'game_info.json' File: '%s'", gameUnmarshallErr.Error())
				}

				// OK - Validate if all files are reachable!
				log.Printf("Successful in reading Campaign save dir: '%s'", currentCampaignPath)
				potentialGame.DataDir = dirName
				return &potentialGame, nil
			}
		}
	} else {
		return nil, fmt.Errorf("could not read Campaign save dir: %s", campaignReadError.Error())
	}
	return nil, fmt.Errorf("could not load campaign by dirName: '%s', no 'game_info.json' match found", dirName)
}
