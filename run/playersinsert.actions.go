package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertActions)
}

func PlayersInsertActions(args ...interface{}) {
	player := args[0].(*game.Player)
	actions, ok := player.Parts[part.Actions].(game.Actions)

	if !ok {
		return
	}

	err := query.ActionsPlayersInsert(player, actions)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("actions insert player")
	}
}
