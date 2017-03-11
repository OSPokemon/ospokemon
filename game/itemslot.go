package game

import (
	"github.com/ospokemon/ospokemon/json"
	"github.com/ospokemon/ospokemon/part"
)

type Itemslot struct {
	Id     int
	Item   *Item
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
	itemslot.Item = item
	itemslot.Amount = amount
	itemslot.AddPart(BuildImaging(item.Animations))
	return itemslot
}

func (i *Itemslot) Part() string {
	return part.Itemslot
}

func (itemslot *Itemslot) Json() json.Json {
	json := json.Json{
		"item":   itemslot.Item.Json(),
		"amount": itemslot.Amount,
	}

	if imaging, _ := itemslot.Parts[part.Imaging].(*Imaging); imaging != nil {
		json["imaging"] = imaging.Json()
	}
	if entity, _ := itemslot.Parts[part.Entity].(*Entity); entity != nil {
		json["entity"] = entity.Id
	}
	if binding, _ := itemslot.Parts[part.Binding].(*Binding); binding != nil {
		json["binding"] = binding.Key
	}

	return json
}

// func (i *Itemslot) Update(u *Universe, e *Entity, d time.Duration) {
// }
