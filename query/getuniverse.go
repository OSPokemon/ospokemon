package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetUniverse(id uint) (*game.Universe, error) {
	if universe, ok := game.Multiverse[id]; ok {
		return universe, nil
	}

	return UniversesSelect(id)
}
