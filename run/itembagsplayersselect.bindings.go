package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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
		log.Add("Username", "2").Add("Error", err.Error()).Error("itembags players select bindings")
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
		log.Add("Username", "2").Add("Bindings", bquery).Warn("itembags players select bindings")
	}
}
