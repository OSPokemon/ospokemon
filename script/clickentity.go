package script

import (
	"ospokemon.com"
)

func ClickEntity(e *ospokemon.Entity, data map[string]interface{}) error {
	if entityId, ok := data["entity"].(float64); ok {
		universe := ospokemon.Multiverse[e.UniverseId]
		entity := universe.Entities[uint(entityId)]

		if dialog := entity.GetDialog(); dialog != nil {
			e.AddPart(dialog)
		}
	}

	return nil
}
