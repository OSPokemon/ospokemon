package engine

import (
	"github.com/ospokemon/ospokemon/space"
	"time"
)

const COMP_Location = "engine/Location"

type Location struct {
	UniverseId uint
	space.Shape
}

func (l *Location) Id() string {
	return COMP_Location
}

func (l *Location) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"universe": l.UniverseId,
		"shape":    l.Shape.Snapshot(),
	}
}

func (l *Location) Update(u *Universe, e *Entity, d time.Duration) {
	// TODO
}
