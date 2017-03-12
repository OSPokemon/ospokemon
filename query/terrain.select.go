package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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

		log.Add("Terrain", id).Info("terrain select")
	}

	return t, err
}
