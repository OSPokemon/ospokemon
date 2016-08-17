package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

type Entity struct {
	Id         uint
	Components map[string]Component
	util.Eventer
}

func (e *Entity) Component(c Component) {
	id := c.Id()

	if e.Components[id] != nil {
		util.Log.Error("Duplicate component id: ", id)
	}

	e.Components[id] = c
}

func (e *Entity) Update(u *Universe, d time.Duration) {
	for _, c := range e.Components {
		c.Update(u, e, d)
	}
}
