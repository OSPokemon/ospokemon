package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectActions)
}

func PlayersSelectActions(args ...interface{}) {
	player := args[0].(*game.Player)

	actions, err := query.ActionsPlayersSelect(player)

	if err != nil {
		log.Add("Username", "2").Add("Error", err.Error()).Error("actions select player")
		return
	}

	player.AddPart(actions)
}
