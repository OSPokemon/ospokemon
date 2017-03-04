package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
)

func init() {
	event.On(event.BindingUp, BindingUpMovement)
}

func BindingUpMovement(args ...interface{}) {
	p := args[0].(*game.Player)
	b := args[1].(*game.Binding)

	if walk, ok := b.Parts[part.Walk].(game.Walk); ok {
		p.Parts[part.Movement].(*game.Movement).ClearWalk(string(walk))
	}
}
