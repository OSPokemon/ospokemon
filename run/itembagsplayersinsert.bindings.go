package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ItembagsPlayersInsert, ItembagsPlayersInsertBindings)
}

func ItembagsPlayersInsertBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag := args[1].(*game.Itembag)

	insert := make(map[string]int)
	for id, itemslot := range itembag.Slots {
		if itemslot == nil {
			continue
		}

		if binding := itemslot.GetBinding(); binding != nil {
			insert[binding.Key] = id
		}
	}

	err := query.BindingsItemsPlayersInsert(player, insert)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("itembags players insert bindings")
	}
}
