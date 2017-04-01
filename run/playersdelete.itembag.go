package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteItembag)
}

func PlayersDeleteItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.ItembagsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("players delete itembags")
	}
}
