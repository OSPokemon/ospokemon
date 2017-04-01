package query

import (
	"ospokemon.com/game"
)

func GetClass(id uint) (*game.Class, error) {
	if class, ok := game.Classes[id]; ok {
		return class, nil
	}
	return ClassesSelect(id)
}
