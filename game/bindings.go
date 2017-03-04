package game

import (
	// "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Bindings map[string]*Binding

func MakeBindingsDefault() Bindings {
	return Bindings{
		"w": &Binding{Key: "w", SystemId: "walk-up", Parts: make(part.Parts)},
		"d": &Binding{Key: "d", SystemId: "walk-right", Parts: make(part.Parts)},
		"a": &Binding{Key: "a", SystemId: "walk-left", Parts: make(part.Parts)},
		"s": &Binding{Key: "s", SystemId: "walk-down", Parts: make(part.Parts)},
	}
}

func (b Bindings) Part() string {
	return part.Bindings
}

func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
	for _, binding := range b {
		binding.Update(u, e, d)
	}
}
