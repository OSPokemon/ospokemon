package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertActions)
}

func PlayersInsertActions(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	actions := player.GetActions()

	err := query.ActionsPlayersInsert(player, actions)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions insert player")
	}
}
