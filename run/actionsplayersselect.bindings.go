package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ActionsPlayersSelect, ActionsPlayersSelectBindings)
}

func ActionsPlayersSelectBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	actions := args[1].(game.Actions)

	aquery, err := query.ActionsBindingsPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("actions players select bindings")
		return
	}

	bindings := player.GetBindings()

	if aquery != nil {
		for key, actionid := range aquery {
			binding := game.MakeBinding()
			binding.Key = key
			action := actions[actionid]

			binding.AddPart(action)
			binding.AddPart(action.GetImaging())

			bindings[key] = binding
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Bindings": bindings,
		}).Warn("actions players select bindings")
	}
}
