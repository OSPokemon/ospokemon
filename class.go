package ospokemon

import "github.com/ospokemon/ospokemon/space"

type Class struct {
	Id         uint
	Dimension  space.Vector
	Animations map[string]string
}

func MakeClass(id uint) *Class {
	c := &Class{
		Id:         id,
		Dimension:  space.Vector{},
		Animations: make(map[string]string),
	}

	return c
}

func GetClass(id uint) (c *Class, err error) {
	if c = Classes.Cache[id]; c == nil {
		if c, err = Classes.Select(id); err == nil {
			Classes.Cache[id] = c
		}
	}
	return
}

// persistence headers
var Classes = struct {
	Cache  map[uint]*Class
	Select func(uint) (*Class, error)
}{make(map[uint]*Class), nil}
