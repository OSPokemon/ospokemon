package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Stats map[string]*Stat

func (s Stats) Part() string {
	return part.Stats
}
