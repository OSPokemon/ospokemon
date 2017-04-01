package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectActions)
}

func PlayersSelectActions(args ...interface{}) {
	player := args[0].(*ospokemon.Player)

	actions, err := query.ActionsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions select player")
		return
	}

	player.AddPart(actions)
}
