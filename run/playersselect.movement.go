package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectMovement)
}

func PlayersSelectMovement(args ...interface{}) {
	p := args[0].(*game.Player)
	m := &game.Movement{}
	p.AddPart(m)
}
