package db

import (
	"github.com/ospokemon/api-go"
)

func LoadSpecies(speciesId int) (ospokemon.Species, error) {
	row := Connection.QueryRow("SELECT id, name, tag, description, xpyield, xpcurve FROM species WHERE id=?", speciesId)

	species := &ospokemon.BasicSpecies{
		STATS: make(map[string]int),
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
		var value int
		rows.Scan(&stat, &value)
		species.STATS[stat] = value
	}

	return species, nil
}
