package game

import (
	"github.com/ospokemon/ospokemon/part"
)

type Item struct {
	Spell
	Tradable bool
	Stack    int
	Value    uint
}

var Items = make(map[uint]*Item)

func MakeItem() *Item {
	return &Item{
		Spell: *MakeSpell(),
	}
}

func (i *Item) Part() string {
	return part.Item
}
