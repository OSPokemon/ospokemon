package script

import (
	"github.com/ospokemon/ospokemon"
)

func ClickEntity(e *ospokemon.Entity, data map[string]interface{}) error {
	if entityId, ok := data["entity"].(float64); ok {
		universe := ospokemon.Universes.Cache[e.UniverseId]
		entity := universe.Entities[uint(entityId)]

		if dialog := entity.GetDialog(); dialog != nil {
			e.AddPart(dialog)
		}
	}

	return nil
}
