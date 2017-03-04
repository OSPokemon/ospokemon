package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.BindingDown, BindingDownAction)
}

func BindingDownAction(args ...interface{}) {
	player := args[0].(*game.Player)
	binding := args[1].(*game.Binding)
	entity := player.Parts[part.Entity].(*game.Entity)
	universe, _ := query.GetUniverse(entity.UniverseId)

	if action, ok := binding.Parts[part.Action].(*game.Action); ok {
		event.Fire(event.ActionCast, universe, entity, action)
	}
}
