package save

import (
	"github.com/ospokemon/ospokemon/space"
)

func NewLocation(shape space.Shape) *Location {
	return &Location{
		UniverseId: 0,
		Shape:      shape,
	}
}
