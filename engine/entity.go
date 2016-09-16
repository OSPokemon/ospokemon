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

func (e *Entity) Snapshot() map[string]interface{} {
	c := make(map[string]interface{})
	m := map[string]interface{}{
		"id":    e.Id,
		"image": e.Image,
		"comp":  c,
	}

	for key, comp := range e.Components {
		if compsnap := comp.Snapshot(); compsnap != nil {
			c[key] = compsnap
		}
	}

	return m
}
