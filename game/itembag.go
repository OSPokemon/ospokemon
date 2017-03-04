package game

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Itembag struct {
	Timers map[uint]*time.Duration
	Slots  []*Itemslot
}

func MakeItembag(size uint) *Itembag {
	bag := &Itembag{
		Timers: make(map[uint]*time.Duration),
		Slots:  make([]*Itemslot, size),
	}

	return bag
}

func (itembag *Itembag) Part() string {
	return part.Itembag
}

func (itembag *Itembag) GetItems() map[uint]int {
	items := make(map[uint]int)

	for _, itemslot := range itembag.Slots {
		if itemslot != nil {
			items[itemslot.Item] = items[itemslot.Item] + itemslot.Amount
		}
	}

	return items
}

func (itembag *Itembag) GetItemslots(itemid uint) []*Itemslot {
	itemslots := make([]*Itemslot, 0)

	for _, itemslot := range itembag.Slots {
		if itemslot != nil && itemslot.Item == itemid {
			itemslots = append(itemslots, itemslot)
		}
	}

	return itemslots
}
