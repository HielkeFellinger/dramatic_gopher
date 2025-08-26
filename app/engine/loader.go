package engine

import (
	"log"
	"os"
	"path/filepath"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"gopkg.in/yaml.v3"
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

			// Test if there is a game_info_file
			currentCampaignPath := filepath.Join(campaignSaveDir, entry.Name())
			if campaignFiles, campaignReadError := os.ReadDir(currentCampaignPath); campaignReadError == nil {
				for _, campaignEntry := range campaignFiles {
					if campaignEntry.IsDir() {
						continue
					}

					// A file has been found; attempt to load it!
					if campaignEntry.Name() == "game_info.yaml" {
						// Attempt to load basics:
						gameInfoData, gameReadErr := os.ReadFile(filepath.Join(currentCampaignPath, campaignEntry.Name()))
						if gameReadErr != nil {
							log.Println("Error Reading campaign 'game_info.yaml' File: ", gameReadErr.Error())
							continue
						}

						var potentialGame BaseGame
						log.Println(string(gameInfoData))
						if gameUnmarshallErr := yaml.Unmarshal(gameInfoData, &potentialGame); gameUnmarshallErr != nil {
							log.Println("Error parsing campaign 'game_info.yaml' File: ", gameUnmarshallErr.Error())
							continue
						}

						// Validate if all files are reachable!
						possibleGames = append(possibleGames, &potentialGame)
					}
				}
			} else {
				log.Println("Could not read campaign saves dir: ", campaignReadError.Error())
			}
		}
	}

	return possibleGames
}
