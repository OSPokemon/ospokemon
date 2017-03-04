package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMovementBindings)
}

func PlayersSelectMovementBindings(args ...interface{}) {
	player := args[0].(*game.Player)

	mquery, err := query.BindingsMovementsPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("players select movement bindings")
		return
	}

	if player.Parts[part.Bindings] == nil {
		player.AddPart(make(game.Bindings))
	}

	if mquery != nil {
		for key, direction := range mquery {
			binding := game.MakeBinding()
			binding.Key = key

			binding.AddPart(game.Walk(direction))

			imaging := game.MakeImaging()
			imaging.Image = "/img/ui/walk/" + direction + ".png"
			binding.AddPart(imaging)

			player.Parts[part.Bindings].(game.Bindings)[key] = binding
		}
	}
}
