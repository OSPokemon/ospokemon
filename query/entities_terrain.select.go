package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func EntitiesTerrainsSelect(universe *game.Universe) (map[uint]*game.Terrain, error) {
	rows, err := Connection.Query(
		"SELECT entity, terrain FROM entities_terrains WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	terrains := make(map[uint]*game.Terrain)

	for rows.Next() {
		var entityId, terrainbuff uint
		err = rows.Scan(&entityId, &terrainbuff)

		if err != nil {
			return nil, err
		}

		terrain, err := GetTerrain(terrainbuff)
		if err != nil {
			return nil, err
		}

		terrains[entityId] = terrain
	}

	log.Add("Universe", universe.Id).Add("Terrains", terrains).Debug("entities_terrains select")

	return terrains, err
}
