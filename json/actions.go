package json

import (
	"github.com/ospokemon/ospokemon/game"
	"strconv"
)

func Actions(actions game.Actions) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	for spellId, action := range actions {
		key := strconv.Itoa(int(spellId))
		_, actionData := Action(action)
		data[key] = actionData
	}

	return "actions", data
}
