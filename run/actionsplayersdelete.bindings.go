package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ActionsPlayersDelete, ActionsPlayersDeleteBindings)
}

func ActionsPlayersDeleteBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	err := query.ActionsBindingsPlayersDelete(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("actions players delete bindings")
	}
}
