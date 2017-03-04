package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Terrain struct {
	Id        uint
	Collision bool
	Image     string
}

var Terrains = make(map[uint]*Terrain)

type TerrainLink uint

func MakeTerrain(id uint) *Terrain {
	terrain := &Terrain{
		Id: id,
	}

	return terrain
}

func (t *Terrain) Part() string {
	return part.Terrain
}
