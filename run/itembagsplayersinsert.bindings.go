package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
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
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("itembags players insert bindings")
	}
}
