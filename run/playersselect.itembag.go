package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectItembag)
}

func PlayersSelectItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag, err := query.ItembagsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("player build itembag")
		return
	}

	player.AddPart(itembag)
}
