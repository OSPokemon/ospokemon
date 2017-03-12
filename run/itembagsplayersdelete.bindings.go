package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ItembagsPlayersDelete, ItembagsPlayersDeleteBindings)
}

func ItembagsPlayersDeleteBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.BindingsItemsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("itembags players delete bindings")
	}
}
