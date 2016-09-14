package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

type Entity struct {
	Id    uint
	Image string
	Components
	util.Eventer
}

func MakeEntity() *Entity {
	return &Entity{
		Components: make(Components),
		Eventer:    make(util.Eventer),
	}
}

func (e *Entity) Update(u *Universe, d time.Duration) {
	for _, c := range e.Components {
		c.Update(u, e, d)
	}
}
