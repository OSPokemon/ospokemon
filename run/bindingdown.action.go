package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/query"
)

func init() {
	event.On(event.BindingDown, BindingDownAction)
}

func BindingDownAction(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)
	entity := player.GetEntity()
	universe, _ := query.GetUniverse(entity.UniverseId)

	if action := binding.GetAction(); action != nil {
		event.Fire(event.ActionCast, universe, entity, action)
	}
}
