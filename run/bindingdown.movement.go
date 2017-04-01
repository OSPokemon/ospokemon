package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
)

func init() {
	event.On(event.BindingDown, BindingDownMovement)
}

func BindingDownMovement(args ...interface{}) {
	player := args[0].(*game.Player)
	bindings := args[1].(*game.Binding)

	movement := player.GetMovement()

	if walk := bindings.GetWalk(); walk != "" {
		movement.Target = nil
		movement.Walk(string(walk))
	}
}
