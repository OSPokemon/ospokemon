package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Itemslot struct {
	Id     int
	Item   uint
	Amount int
	part.Parts
}

type Itemslots []*Itemslot

func MakeItemslot() *Itemslot {
	itemslot := &Itemslot{
		Id:    -1,
		Parts: make(part.Parts),
	}

	return itemslot
}

func (i *Itemslot) Part() string {
	return part.Itemslot
}

func (itemslots Itemslots) Part() string {
	return part.Itemslots
}

// func (i *Itemslot) Update(u *Universe, e *Entity, d time.Duration) {
// }
