package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.ItembagsPlayersDelete, ItembagsPlayersDeleteBindings)
}

func ItembagsPlayersDeleteBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	err := persistence.BindingsItemsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("itembags players delete bindings")
	}
}
