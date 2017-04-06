package ospokemon

import (
	"ospokemon.com/json"
	"time"
)

const PARTitembag = "itembag"

type Itembag struct {
	Timers map[uint]*time.Duration
	Slots  map[uint]*Itemslot
}

func MakeItembag() *Itembag {
	bag := &Itembag{
		Timers: make(map[uint]*time.Duration),
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

func (itembag *Itembag) Json() json.Json {
	data := json.Json{}
	for id, itemslot := range itembag.Slots {
		if itemslot == nil {
			data[json.StringUint(id)] = nil
		} else {
			itemslotJson := itemslot.Json()
			itemslotJson["timer"] = json.FmtDuration(itembag.Timers[itemslot.Item.Id])
			data[json.StringUint(id)] = itemslotJson
		}
	}
	return data
}
