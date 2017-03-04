package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetType(id uint) (*game.Type, error) {
	if t, ok := game.Types[id]; ok {
		return t, nil
	}

	return TypesSelect(id)
}
