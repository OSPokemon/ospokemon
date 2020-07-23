package run

import (
	"math"
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.Movement, MovementImaging)
}

func MovementImaging(args ...interface{}) {
	e := args[0].(*ospokemon.Entity)
	v, _ := args[1].(space.Vector)

	i := e.GetImaging()

	if v.DX == 0 && v.DY == 0 {
		i.Image = i.Animations["portrait"]
	} else if slope := v.AsSlope(); slope == math.Inf(-1) {
		i.Image = i.Animations["walk-up"]
	} else if slope == math.Inf(1) {
		i.Image = i.Animations["walk-down"]
	} else if v.DX > 0 {
		i.Image = i.Animations["walk-right"]
	} else {
		i.Image = i.Animations["walk-left"]
	}
}
