package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersDelete, PlayersDeleteActions)
}

func PlayersDeleteActions(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.ActionsPlayersDelete(player)

	if err != nil {
		logrus.Error(err.Error())
	}
}
