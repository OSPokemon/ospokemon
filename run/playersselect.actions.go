package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectActions)
}

func PlayersSelectActions(args ...interface{}) {
	player := args[0].(*game.Player)

	actions, err := query.ActionsPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("actions select player")
		return
	}

	player.AddPart(actions)
}
