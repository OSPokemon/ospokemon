package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetSpecies(id uint) (*game.Species, error) {
	if species, ok := game.Specieses[id]; ok {
		return species, nil
	}

	return SpeciesSelect(id)
}
