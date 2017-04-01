package ospokemon

import (
	"ospokemon.com/space"
)

type Class struct {
	Id         uint
	Dimension  space.Vector
	Animations map[string]string
}

var Classes = make(map[uint]*Class)

func MakeClass(id uint) *Class {
	c := &Class{
		Id:         id,
		Dimension:  space.Vector{},
		Animations: make(map[string]string),
	}

	return c
}
