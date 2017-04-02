package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteItembag)
}

func PlayersDeleteItembag(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	err := persistence.ItembagsPlayersDelete(player)

	if err != nil {
		log.Add("Player", player.Username).Add("Error", err.Error()).Error("players delete itembags")
	}
}
