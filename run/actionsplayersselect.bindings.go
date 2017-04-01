package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.ActionsPlayersSelect, ActionsPlayersSelectBindings)
}

func ActionsPlayersSelectBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	actions := args[1].(ospokemon.Actions)

	aquery, err := query.ActionsBindingsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions players select bindings")
		return
	}

	bindings := player.GetBindings()

	if aquery != nil {
		for key, actionid := range aquery {
			binding := ospokemon.MakeBinding()
			binding.Key = key
			action := actions[actionid]

			binding.AddPart(action)
			binding.AddPart(action.GetImaging())

			bindings[key] = binding
		}
	} else {
		log.Add("Username", player.Username).Add("Bindings", bindings).Warn("actions players select bindings")
	}
}
