package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ItembagsPlayersSelect, ItembagsPlayersSelectBindings)
}

func ItembagsPlayersSelectBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag := args[1].(*game.Itembag)

	bquery, err := query.BindingsItemsPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("itembags players select bindings")
		return
	}

	bindings := player.GetBindings()

	if bquery != nil {
		for key, itemslotid := range bquery {
			binding := game.MakeBinding()
			binding.Key = key
			itemslot := itembag.Slots[itemslotid]
			itemslot.AddPart(binding)
			binding.Parts = itemslot.Parts

			bindings[key] = binding
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Bindings": bquery,
		}).Warn("itembags players select bindings")
	}
}
