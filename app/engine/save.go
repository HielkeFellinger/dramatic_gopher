package engine

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/HielkeFellinger/dramatic_gopher/app/ecs"
	"github.com/HielkeFellinger/dramatic_gopher/app/ecs/core"
)

func SaveBaseGame(game *BaseGame) error {

	world := *game.World
	rawGame := RawSaveFile{}

	// Override time
	rawGame.DateTime = time.Now().Format("2006-01-02_15:04:05_MST")

	// Split per type;
	log.Println("   - Saving Item (Entities)")
	rawGame.Items = parseEntityIntoRawEntity(world.GetItemEntities())
	log.Println("   - Saving Character (Entities)")
	rawGame.Characters = parseEntityIntoRawEntity(world.GetCharacterEntities())
	log.Println("   - Saving Map (Entities)")
	rawGame.Maps = parseEntityIntoRawEntity(world.GetMapEntities())
	log.Println("   - Saving Map Content (Entities)")
	rawGame.Storage = parseEntityIntoRawEntity(world.GetStorageEntities())

	// Marshal the game data
	gameFileContent, err := json.Marshal(rawGame)
	if err != nil {
		return err
	}

	// Attempt to save the game
	log.Printf("Attempting to save game '%s/%s' to file: '%s'", game.Id, game.Title, game.SaveFile)
	if writeErr := os.WriteFile(game.SaveFile, gameFileContent, 0644); writeErr != nil {
		return writeErr
	}
	log.Println("Saved the game")

	return nil
}

func parseEntityIntoRawEntity(entities []ecs.Entity) []core.RawEntity {
	rawEntities := make([]core.RawEntity, 0)

	for _, entity := range entities {
		// Skip nil ptr; may be a leftover of slices.delete not reducing the total size of the underlying array.
		if entity == nil {
			continue
		}

		rawEntity := core.RawEntity{
			Id:         entity.GetId().String(),
			Components: parseComponentsToRawComponents(entity.GetComponents()),
		}
		rawEntities = append(rawEntities, rawEntity)
	}
	return rawEntities
}

func parseComponentsToRawComponents(components []*core.Component) []core.RawComponent {
	rawComponents := make([]core.RawComponent, 0)

	for _, component := range components {
		// Skip nil ptr; may be a leftover of slices.delete not reducing the total size of the underlying array.
		if component == nil {
			continue
		}

		if rawComponent, err := (*component).ParseToRawComponent(); err == nil {
			rawComponents = append(rawComponents, rawComponent)
		} else {
			log.Printf("Error parsing component of type: '%v' with error message: '%s'", component, err.Error())
		}
	}
	return rawComponents
}
