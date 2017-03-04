package json

import (
	"github.com/ospokemon/ospokemon/game"
)

func Player(p *game.Player) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"username": p.Username,
		"level":    p.Level,
	}

	// if expand {
	// 	for _, part := range p.Parts {
	// 		if jsoner, ok := part.(Jsoner); ok {
	// 			key, partData := jsoner.Json(false)
	// 			data[key] = partData
	// 		}
	// 	}
	// }

	return "player", data
}
