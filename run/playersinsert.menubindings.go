package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertMenuBindings)
}

func PlayersInsertMenuBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	bindings := player.GetBindings()

	if len(bindings) < 1 {
		bindings = make(game.Bindings)

		menubindings := map[string]game.Menu{
			"Enter":  "chat",
			"c":      "player",
			"b":      "itembag",
			"x":      "actions",
			"Escape": "settings",
		}

		for key, menu := range menubindings {
			binding := game.MakeBinding()
			binding.Key = key
			binding.AddPart(menu)
			bindings[binding.Key] = binding
		}

		log.Add("Username", player.Username).Debug("players insert menubindings: grant default menu bindings")
	}

	insert := make(map[string]string)
	for key, binding := range bindings {
		if menu := binding.GetMenu(); menu != "" {
			insert[key] = string(menu)
		}
	}

	err := query.BindingsMenusPlayersInsert(player, insert)

	if err != nil {
		log.Add("Username", player.Username).Add("Bindings", bindings).Add("Error", err.Error()).Error("players insert menubindings")
	}
}
