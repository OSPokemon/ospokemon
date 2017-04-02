package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertActions)
}

func PlayersInsertActions(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	actions := player.GetActions()

	err := persistence.ActionsPlayersInsert(player, actions)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("actions insert player")
	}
}
