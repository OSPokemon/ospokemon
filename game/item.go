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

func MakeItem(id uint) *Item {
	i := &Item{
		Spell: Spell{
			Id:         id,
			Animations: make(map[string]string),
			Data:       make(map[string]string),
		},
	}

	return i
}

func (i *Item) Part() string {
	return part.Item
}
