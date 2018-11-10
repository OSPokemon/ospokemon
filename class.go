package ospokemon

import "ospokemon.com/space"

type Class struct {
	Id         uint
	Dimension  space.Vector
	Animations map[string]string
}

var classes = make(map[uint]*Class)

func MakeClass(id uint) *Class {
	c := &Class{
		Id:         id,
		Dimension:  space.Vector{},
		Animations: make(map[string]string),
	}

	return c
}

func GetClass(id uint) (*Class, error) {
	if classes[id] == nil {
		if c, err := Classes.Select(id); err == nil {
			classes[id] = c
		} else {
			return nil, err
		}
	}

	return classes[id], nil
}

// persistence headers
var Classes struct {
	Select func(uint) (*Class, error)
}
