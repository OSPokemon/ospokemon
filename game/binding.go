package game

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Binding struct {
	Key      string
	SystemId string
	part.Parts
}

func MakeBinding() *Binding {
	return &Binding{
		Parts: make(part.Parts),
	}
}

func (b *Binding) Update(u *Universe, e *Entity, d time.Duration) {
}
