package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func EntitiesTerrainSelect(entity *game.Entity, universe *game.Universe) (*game.Terrain, error) {
	row := Connection.QueryRow(
		"SELECT terrain FROM entities_terrain WHERE entity=? AND universe=?",
		entity.Id,
		universe.Id,
	)

	var terrainbuff uint
	var terrain *game.Terrain
	err := row.Scan(&terrainbuff)

	if err == nil {
		terrain, err = GetTerrain(terrainbuff)

		if err == nil {
			log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Terrain", terrain.Id).Debug("entities_terrain select")
		}
	}

	return terrain, err
}
