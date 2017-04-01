package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
)

func init() {
	event.On(event.BindingUp, BindingUpMovement)
}

func BindingUpMovement(args ...interface{}) {
	p := args[0].(*ospokemon.Player)
	b := args[1].(*ospokemon.Binding)

	if walk := b.GetWalk(); walk != "" {
		p.GetMovement().ClearWalk(string(walk))
	}
}
