package query

import (
	"ospokemon.com"
)

func GetClass(id uint) (*ospokemon.Class, error) {
	if class, ok := ospokemon.Classes[id]; ok {
		return class, nil
	}
	return ClassesSelect(id)
}
