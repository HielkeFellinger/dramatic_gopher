package engine

import (
	"log"
	"os"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
)

func FindAvailableGames() []string {
	possibleGames := make([]string, 0)

	entries, readErr := os.ReadDir(config.CurrentConfig.CampaignSavesDir)
	if readErr != nil {
		log.Println("Could not read campaign saves dir: ", readErr)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			possibleGames = append(possibleGames, entry.Name())
		}
	}

	return possibleGames
}
