package query

import (
	"ospokemon.com"
)

func GetType(id uint) (*ospokemon.Type, error) {
	if t, ok := ospokemon.Types[id]; ok {
		return t, nil
	}

	return TypesSelect(id)
}
