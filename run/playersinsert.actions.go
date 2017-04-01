package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertActions)
}

func PlayersInsertActions(args ...interface{}) {
	player := args[0].(*game.Player)
	actions := player.GetActions()

	err := query.ActionsPlayersInsert(player, actions)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions insert player")
	}
}
