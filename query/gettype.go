package query

import (
	"ospokemon.com/game"
)

func GetType(id uint) (*game.Type, error) {
	if t, ok := game.Types[id]; ok {
		return t, nil
	}

	return TypesSelect(id)
}
