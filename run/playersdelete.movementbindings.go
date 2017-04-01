package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteMovementBindings)
}

func PlayersDeleteMovementBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	err := query.BindingsMovementsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("players delete movementbindings")
	}
}
