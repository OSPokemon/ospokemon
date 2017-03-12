package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertActions)
}

func PlayersInsertActions(args ...interface{}) {
	player := args[0].(*game.Player)
	actions := player.GetActions()

	err := query.ActionsPlayersInsert(player, actions)

	if err != nil {
		log.Add("Username", "2").Add("Error", err.Error()).Error("actions insert player")
	}
}
