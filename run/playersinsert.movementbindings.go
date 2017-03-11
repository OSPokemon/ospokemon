package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertMovementBindings)
}

func PlayersInsertMovementBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	bindings, ok := player.Parts[part.Bindings].(game.Bindings)

	if !ok {
		bindings = make(game.Bindings)

		movementbindings := map[string]game.Walk{
			"a": "left",
			"s": "down",
			"d": "right",
			"w": "up",
		}

		for key, direction := range movementbindings {
			binding := game.MakeBinding()
			binding.Key = key
			binding.AddPart(direction)
			bindings[binding.Key] = binding
		}

		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players insert movement bindings: grant default movement bindings")
	}

	insert := make(map[string]string)
	for key, binding := range bindings {
		if walk, ok := binding.Parts[part.Walk].(game.Walk); ok {
			insert[key] = string(walk)
		}
	}

	err := query.BindingsMovementsPlayersInsert(player, insert)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Bindings": bindings,
			"Error":    err.Error(),
		}).Error("players insert movement bindings")
	}
}