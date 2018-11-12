package persistence

import (
	"github.com/pkg/errors"
	"ospokemon.com"
	"ztaylor.me/log"
)

func EntitiesTerrainsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Terrain, error) {
	rows, err := Connection.Query(
		"SELECT entity, terrain FROM entities_terrains WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "entities_terrains.select")
	}

	buf := make(map[uint]uint)

	for rows.Next() {
		var entityId, terrainbuff uint
		err = rows.Scan(&entityId, &terrainbuff)

		if err != nil {
			return nil, errors.Wrap(err, "entities_terrains.scan")
		}

		buf[entityId] = terrainbuff
	}
	rows.Close()

	terrains := make(map[uint]*ospokemon.Terrain)

	for k, v := range buf {
		terrain, err := ospokemon.GetTerrain(v)
		if err != nil {
			return nil, err
		}

		terrains[k] = terrain
	}

	log.Add("Universe", universe.Id).Add("Terrains", buf).Debug("entities_terrains select")

	return terrains, err
}
