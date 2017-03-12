package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
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
