package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func EntitiesTerrainsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Terrain, error) {
	rows, err := Connection.Query(
		"SELECT entity, terrain FROM entities_terrains WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	terrains := make(map[uint]*ospokemon.Terrain)

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
