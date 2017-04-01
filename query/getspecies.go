package query

import (
	"ospokemon.com"
)

func GetSpecies(id uint) (*ospokemon.Species, error) {
	if species, ok := ospokemon.Specieses[id]; ok {
		return species, nil
	}

	return SpeciesSelect(id)
}
