package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertMovementBindings)
}

func PlayersInsertMovementBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	bindings := player.GetBindings()

	if len(bindings) < 1 {
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

		log.Add("Username", "2").Debug("players insert movement bindings: grant default movement bindings")
	}

	insert := make(map[string]string)
	for key, binding := range bindings {
		if walk := binding.GetWalk(); walk != "" {
			insert[key] = string(walk)
		}
	}

	err := query.BindingsMovementsPlayersInsert(player, insert)

	if err != nil {
		log.Add("Username", player.Username).Add("Bindings", bindings).Add("Error", err.Error()).Error("players insert movement bindings")
	}
}
