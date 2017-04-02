package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectItembag)
}

func PlayersSelectItembag(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	itembag, err := persistence.ItembagsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("player build itembag")
		return
	}

	player.AddPart(itembag)
}
