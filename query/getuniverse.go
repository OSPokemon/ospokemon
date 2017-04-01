package query

import (
	"ospokemon.com/game"
)

func GetUniverse(id uint) (*game.Universe, error) {
	if universe, ok := game.Multiverse[id]; ok {
		return universe, nil
	}

	return UniversesSelect(id)
}
