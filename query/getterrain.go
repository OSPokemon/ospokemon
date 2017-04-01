package query

import (
	"ospokemon.com/game"
)

func GetTerrain(id uint) (*game.Terrain, error) {
	if t, ok := game.Terrains[id]; ok {
		return t, nil
	}

	return TerrainSelect(id)
}
