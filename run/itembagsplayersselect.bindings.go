package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.ItembagsPlayersSelect, ItembagsPlayersSelectBindings)
}

func ItembagsPlayersSelectBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	itembag := args[1].(*ospokemon.Itembag)

	bquery, err := query.BindingsItemsPlayersSelect(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("itembags players select bindings")
		return
	}

	bindings := player.GetBindings()

	if bquery != nil {
		for key, itemslotid := range bquery {
			binding := ospokemon.MakeBinding()
			binding.Key = key
			itemslot := itembag.Slots[itemslotid]
			itemslot.AddPart(binding)
			binding.Parts = itemslot.Parts

			bindings[key] = binding
		}
	} else {
		log.Add("Username", player.Username).Add("Bindings", bquery).Warn("itembags players select bindings")
	}
}
