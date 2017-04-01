package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func TerrainSelect(id uint) (*ospokemon.Terrain, error) {
	row := Connection.QueryRow(
		"SELECT collision, image FROM terrain WHERE id=?",
		id,
	)

	t := ospokemon.MakeTerrain(id)
	err := row.Scan(&t.Collision, &t.Image)

	if err == nil {
		ospokemon.Terrains[id] = t

		log.Add("Terrain", id).Info("terrain select")
	}

	return t, err
}
