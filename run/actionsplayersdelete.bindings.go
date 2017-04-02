package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.ActionsPlayersDelete, ActionsPlayersDeleteBindings)
}

func ActionsPlayersDeleteBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	err := persistence.ActionsBindingsPlayersDelete(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions players delete bindings")
	}
}
