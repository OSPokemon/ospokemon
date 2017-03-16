package game

import (
	"github.com/ospokemon/ospokemon/space"
)

const PARTitem = "item"

type Item struct {
	Spell
	Dimension space.Vector
	Tradable  bool
	Stack     int
	Value     uint
}

var Items = make(map[uint]*Item)

func MakeItem() *Item {
	return &Item{
		Dimension: space.Vector{},
		Spell:     *MakeSpell(),
	}
}

func (i *Item) Part() string {
	return PARTitem
}

func (parts Parts) GetItem() *Item {
	item, _ := parts[PARTitem].(*Item)
	return item
}
