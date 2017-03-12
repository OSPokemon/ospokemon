package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func init() {
	event.On(event.BindingUp, BindingUpMovement)
}

func BindingUpMovement(args ...interface{}) {
	p := args[0].(*game.Player)
	b := args[1].(*game.Binding)

	if walk := b.GetWalk(); walk != "" {
		p.GetMovement().ClearWalk(string(walk))
	}
}
