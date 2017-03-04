package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func TerrainSelect(id uint) (*game.Terrain, error) {
	row := Connection.QueryRow(
		"SELECT collision, image FROM terrain WHERE id=?",
		id,
	)

	t := game.MakeTerrain(id)
	err := row.Scan(&t.Collision, &t.Image)

	if err == nil {
		game.Terrains[id] = t

		logrus.WithFields(logrus.Fields{
			"Terrain": t.Id,
		}).Info("terrain select")
	}

	return t, err
}
