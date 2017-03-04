package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetClass(id uint) (*game.Class, error) {
	if class, ok := game.Classes[id]; ok {
		return class, nil
	}
	return ClassesSelect(id)
}
