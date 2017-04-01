package query

import (
	"ospokemon.com"
)

func GetTerrain(id uint) (*ospokemon.Terrain, error) {
	if t, ok := ospokemon.Terrains[id]; ok {
		return t, nil
	}

	return TerrainSelect(id)
}
