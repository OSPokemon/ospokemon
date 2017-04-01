package query

import (
	"ospokemon.com/game"
)

func GetSpecies(id uint) (*game.Species, error) {
	if species, ok := game.Specieses[id]; ok {
		return species, nil
	}

	return SpeciesSelect(id)
}
