package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func init() {
	ospokemon.Terrains.Select = TerrainSelect
}

func TerrainSelect(id uint) (*ospokemon.Terrain, error) {
	row := Connection.QueryRow(
		"SELECT collision, image FROM terrain WHERE id=?",
		id,
	)

	t := ospokemon.MakeTerrain(id)
	err := row.Scan(&t.Collision, &t.Image)

	if err == nil {
		log.Add("Terrain", id).Info("terrain select")
	}

	return t, err
}
