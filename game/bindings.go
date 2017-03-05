package game

import (
	"github.com/ospokemon/ospokemon/part"
	// "time"
)

type Bindings map[string]*Binding

func (b Bindings) Part() string {
	return part.Bindings
}

// func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
// 	for _, binding := range b {
// 		binding.Update(u, e, d)
// 	}
// }
