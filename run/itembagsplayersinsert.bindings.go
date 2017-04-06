package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.ItembagsPlayersInsert, ItembagsPlayersInsertBindings)
}

func ItembagsPlayersInsertBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	itembag := args[1].(*ospokemon.Itembag)

	insert := make(map[string]uint)
	for id, itemslot := range itembag.Slots {
		if itemslot == nil {
			continue
		}

		if binding := itemslot.GetBinding(); binding != nil {
			insert[binding.Key] = id
		}
	}

	err := persistence.BindingsItemsPlayersInsert(player, insert)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("itembags players insert bindings")
	}
}
