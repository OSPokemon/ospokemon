package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
)

func init() {
	event.On(event.BindingDown, BindingDownMovement)
}

func BindingDownMovement(args ...interface{}) {
	player := args[0].(*game.Player)
	bindings := args[1].(*game.Binding)

	movement := player.Parts[part.Movement].(*game.Movement)

	if walk, ok := bindings.Parts[part.Walk].(game.Walk); ok {
		movement.Target = nil
		movement.Walk(string(walk))
	}
}
