package engine

import (
	"github.com/ospokemon/ospokemon/space"
	"time"
)

const COMP_Location = "engine/Location"

type Location struct {
	space.Shape
}

func (l *Location) Id() string {
	return COMP_Location
}

func (l *Location) Update(u *Universe, e *Entity, d time.Duration) {
	// TODO
}
