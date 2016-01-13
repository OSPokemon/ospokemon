package db

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/objects"
)

func LoadSpecies(speciesId int) (*objects.Species, error) {
	row := Connection.QueryRow("SELECT id, name, tag, description, xpyield, xpcurve FROM species WHERE id=?", speciesId)

	species := &objects.Species{
		BasicSpecies: ospokemon.MakeBasicSpecies(""),
		GRAPHICS:     make(map[engine.AnimationType]string),
	}

	err := row.Scan(&species.ID, &species.NAME, &species.TAG, &species.DESCRIPTION, &species.EXPERIENCEYIELD, &species.EXPERIENCECURVE)
	if err != nil {
		return nil, err
	}

	rows, err := Connection.Query("SELECT stat, value FROM species_stats WHERE speciesid=?", speciesId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var stat string
		var value float64
		rows.Scan(&stat, &value)
		species.STATS[stat] = value
	}

	rows, err = Connection.Query("SELECT anim, image FROM species_animations WHERE speciesid=?", speciesId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var anim, image string
		rows.Scan(&anim, &image)
		species.GRAPHICS[engine.AnimationType(anim)] = image
	}

	return species, nil
}

func GetSpeciesIds() []int {
	rows, err := Connection.Query("SELECT id FROM species")
	if err != nil {
		return nil
	}

	ids := make([]int, 0)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids
}
