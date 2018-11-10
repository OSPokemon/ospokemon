package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func EntitiesTerrainsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Terrain, error) {
	rows, err := Connection.Query(
		"SELECT entity, terrain FROM entities_terrains WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	terrainslog := make(map[uint]uint)
	terrains := make(map[uint]*ospokemon.Terrain)

	for rows.Next() {
		var entityId, terrainbuff uint
		err = rows.Scan(&entityId, &terrainbuff)

		if err != nil {
			return nil, err
		}

		terrain, err := ospokemon.GetTerrain(terrainbuff)
		if err != nil {
			return nil, err
		}

		terrains[entityId] = terrain
		terrainslog[entityId] = terrain.Id
	}

	log.Add("Universe", universe.Id).Add("Terrains", terrainslog).Debug("entities_terrains select")

	return terrains, err
}
