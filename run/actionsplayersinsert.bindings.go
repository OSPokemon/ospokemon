package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
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
		if bindings, ok := action.Parts[part.Bindings].(game.Bindings); ok {
			for key, _ := range bindings {
				insert[key] = action.Spell
			}
		}
	}

	err := query.ActionsBindingsPlayersInsert(player, insert)

	if err != nil {
		logrus.Error(err.Error())
	}
}
