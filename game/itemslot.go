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

func MakeItemslot() *Itemslot {
	itemslot := &Itemslot{
		Id:    -1,
		Parts: make(part.Parts),
	}

	return itemslot
}

func BuildItemslot(id int, item *Item, amount int) *Itemslot {
	itemslot := MakeItemslot()
	itemslot.Id = id
	itemslot.Item = item.Id
	itemslot.Amount = amount
	itemslot.AddPart(BuildImaging(item.Animations))
	return itemslot
}

func (i *Itemslot) Part() string {
	return part.Itemslot
}

// func (i *Itemslot) Update(u *Universe, e *Entity, d time.Duration) {
// }
