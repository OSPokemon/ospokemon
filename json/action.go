package json

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func Action(a *game.Action) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	if a.Timer != nil {
		data["timer"] = int64(*a.Timer)
	} else {
		data["timer"] = 0
	}

	if spell, _ := query.GetSpell(a.Spell); spell != nil {
		data["spell"] = spell.Snapshot()
	} else {
		data["spell"] = a.Spell
	}

	// if expand {
	// 	for _, part := range a.Parts {
	// 		if jsoner, ok := part.(Jsoner); ok {
	// 			key, jsonerData := jsoner.Json(false)
	// 			data[key] = jsonerData
	// 		}
	// 	}
	// }

	return "action", data
}
