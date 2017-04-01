package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
)

func init() {
	event.On(event.BindingDown, BindingDownMovement)
}

func BindingDownMovement(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	bindings := args[1].(*ospokemon.Binding)

	movement := player.GetMovement()

	if walk := bindings.GetWalk(); walk != "" {
		movement.Target = nil
		movement.Walk(string(walk))
	}
}
