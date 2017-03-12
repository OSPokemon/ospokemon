package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ActionsPlayersInsert, ActionsPlayersInsertBindings)
}

func ActionsPlayersInsertBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	actions := args[1].(game.Actions)

	insert := make(map[string]uint)
	for _, action := range actions {
		if binding := action.GetBinding(); binding != nil {
			insert[binding.Key] = action.Spell.Id
		}
	}

	err := query.ActionsBindingsPlayersInsert(player, insert)

	if err != nil {
		logrus.Error(err.Error())
	}
}
