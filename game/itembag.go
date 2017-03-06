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

func (itembag *Itembag) Add(item *Item, amount int) bool {
	for _, itemslot := range itembag.GetItemslots(item.Id) {
		itemslot.Amount += amount

		if itemslot.Amount < item.Stack {
			return true
		}
		amount = itemslot.Amount - item.Stack
	}

	for id, itemslot := range itembag.Slots {
		if itemslot == nil {
			itemslot = BuildItemslot(id, item, amount)
			itembag.Slots[id] = itemslot

			if itemslot.Amount <= item.Stack {
				return true
			}
			amount = itemslot.Amount - item.Stack
		}
	}

	return false
}

func (itembag *Itembag) Remove(item *Item, amount int) bool {
	if itembag.GetItems()[item.Id] < amount {
		return false
	}

	for _, itemslot := range itembag.GetItemslots(item.Id) {
		itemslot.Amount -= amount
		amount = 0

		if itemslot.Amount < 1 {
			amount -= itemslot.Amount
			itembag.Slots[itemslot.Id] = nil
		}

		if amount < 1 {
			return true
		}
	}

	return false
}
