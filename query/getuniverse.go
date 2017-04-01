package query

import (
	"ospokemon.com"
)

func GetUniverse(id uint) (*ospokemon.Universe, error) {
	if universe, ok := ospokemon.Multiverse[id]; ok {
		return universe, nil
	}

	return UniversesSelect(id)
}
