package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteMovementBindings)
}

func PlayersDeleteMovementBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.BindingsMovementsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("players delete movementbindings")
	}
}
