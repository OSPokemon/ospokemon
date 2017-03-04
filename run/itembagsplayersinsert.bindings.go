package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ItembagsPlayersInsert, ItembagsPlayersInsertBindings)
}

func ItembagsPlayersInsertBindings(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag := args[1].(*game.Itembag)

	insert := make(map[string]uint)
	for itemid, _ := range itembag.GetItems() {
		itemslots := itembag.GetItemslots(itemid)
		itemslot := itemslots[0]

		if bindings, ok := itemslot.Parts[part.Bindings].(game.Bindings); ok {
			for key, _ := range bindings {
				insert[key] = itemslot.Item
			}
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
