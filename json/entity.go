package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Entity(e *game.Entity) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"id":    e.Id,
		"shape": e.Shape.Snapshot(),
	}

	for _, part := range e.Parts {
		if player, ok := part.(*game.Player); ok {
			key, partData := Player(player)
			data[key] = partData
		} else if imaging, ok := part.(*game.Imaging); ok {
			key, partData := Imaging(imaging)
			data[key] = partData
		}
	}

	return "entity", data
}
