package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.ActionsPlayersInsert, ActionsPlayersInsertBindings)
}

func ActionsPlayersInsertBindings(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	actions := args[1].(ospokemon.Actions)

	insert := make(map[string]uint)
	for _, action := range actions {
		if binding := action.GetBinding(); binding != nil {
			insert[binding.Key] = action.Spell.Id
		}
	}

	err := persistence.ActionsBindingsPlayersInsert(player, insert)

	if err != nil {
		log.Add("player", player.Username).Add("Error", err.Error()).Error("actions players insert bindings")
	}
}
