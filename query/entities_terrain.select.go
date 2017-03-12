package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
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
			logrus.WithFields(logrus.Fields{
				"Universe": universe.Id,
				"Entity":   entity.Id,
				"Terrain":  terrain.Id,
			}).Debug("entities_terrain select")
		}
	}

	return terrain, err
}
