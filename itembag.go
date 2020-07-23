package ospokemon

import (
	"time"

	"taylz.io/types"
)

const PARTitembag = "itembag"

type Itembag struct {
	Timers map[uint]*Timer
	Slots  map[uint]*Itemslot
}

func MakeItembag() *Itembag {
	bag := &Itembag{
		Timers: make(map[uint]*Timer),
		Slots:  make(map[uint]*Itemslot),
	}

	return bag
}

func (itembag *Itembag) Part() string {
	return PARTitembag
}

func (parts Parts) GetItembag() *Itembag {
	itembag, _ := parts[PARTitembag].(*Itembag)
	return itembag
}

func (itembag *Itembag) GetItems() map[uint]int {
	items := make(map[uint]int)

	for itemid, itemslot := range itembag.Slots {
		if itemslot != nil {
			items[itemid] = itemslot.Amount
		}
	}

	return items
}

func (itembag *Itembag) Add(item *Item, amount int) bool {
	if itembag.GetItems()[item.Id]+amount > item.Stack {
		return false
	}

	if itemslot := itembag.Slots[item.Id]; itemslot == nil {
		itembag.Slots[item.Id] = BuildItemslot(item, amount)
	} else {
		itemslot.Amount += amount
	}

	return true
}

func (itembag *Itembag) Remove(item *Item, amount int) bool {
	if itembag.GetItems()[item.Id] < amount {
		return false
	}

	itemslot := itembag.Slots[item.Id]
	itemslot.Amount -= amount

	return true
}

func (itembag *Itembag) Json() types.Dict {
	data := types.Dict{}

	for id, itemslot := range itembag.Slots {
		if itemslot == nil {
			data[types.StringUint(id)] = nil
		} else {
			itemslotJson := itemslot.Json()
			itemslotJson["timer"] = itembag.Timers[itemslot.Item.Id].Fmt()
			data[types.StringUint(id)] = itemslotJson
		}
	}

	return data
}

func (itembag *Itembag) Update(u *Universe, e *Entity, d time.Duration) {
	for itemid, timer := range itembag.Timers {
		if timer == nil {
			continue
		}

		item, err := GetItem(itemid)
		if err != nil {
			log.Add("Error", err).Add("ItemId", itemid).Error("itembag: update")
			continue
		}

		td := timer.Duration()

		timer.Set(td - d)

		if td <= item.Cooldown && td+d > item.Cooldown {
			itemslot := itembag.Slots[itemid]
			err := itemslot.Item.Run(e)

			if err != nil {
				log.Add("Error", err).Add("ItemId", itemid).Error("itembag: itemcast")
			}
		}

		if *timer <= 0 {
			itembag.Timers[itemid] = nil
		}
	}
}
