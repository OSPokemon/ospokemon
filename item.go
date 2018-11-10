package ospokemon

import "ospokemon.com/space"

const PARTitem = "item"

type Item struct {
	Spell
	Dimension space.Vector
	Tradable  bool
	Stack     int
	Value     uint
}

var items = make(map[uint]*Item)

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

func GetItem(id uint) (*Item, error) {
	if items[id] == nil {
		if i, err := Items.Select(id); err == nil {
			items[id] = i
		} else {
			return nil, err
		}
	}

	return items[id], nil
}

// persistence headers
var Items struct {
	Select func(uint) (*Item, error)
}
