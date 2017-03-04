package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMenuBindings)
}

func PlayersSelectMenuBindings(args ...interface{}) {
	player := args[0].(*game.Player)

	mquery, err := query.BindingsMenusPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("players select menubindings")
		return
	}

	bindings, ok := player.Parts[part.Bindings].(game.Bindings)
	if !ok {
		bindings = make(game.Bindings)
		player.AddPart(bindings)
	}

	if mquery != nil {
		for key, menu := range mquery {
			binding := game.MakeBinding()
			binding.Key = key

			binding.AddPart(game.Menu(menu))

			imaging := game.MakeImaging()
			imaging.Image = "/img/ui/menu/" + menu + ".png"
			binding.AddPart(imaging)

			bindings[key] = binding
		}
	}
}
