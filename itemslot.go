package ospokemon

import "taylz.io/types"

const PARTitemslot = "itemslot"

type Itemslot struct {
	Item   *Item
	Sort   int
	Amount int
	Parts
}

func MakeItemslot() *Itemslot {
	itemslot := &Itemslot{
		Parts: make(Parts),
	}

	return itemslot
}

func BuildItemslot(item *Item, amount int) *Itemslot {
	itemslot := MakeItemslot()
	itemslot.Item = item
	itemslot.Amount = amount
	itemslot.AddPart(BuildImaging(item.Animations))
	return itemslot
}

func (i *Itemslot) Part() string {
	return PARTitemslot
}

func (parts Parts) GetItemslot() *Itemslot {
	itemslot, _ := parts[PARTitemslot].(*Itemslot)
	return itemslot
}

func (itemslot *Itemslot) Json() types.Dict {
	json := types.Dict{
		"item":   itemslot.Item.Json(),
		"amount": itemslot.Amount,
	}

	if imaging := itemslot.GetImaging(); imaging != nil {
		json["image"] = imaging.Image
	}
	if entity := itemslot.GetEntity(); entity != nil {
		json["entity"] = entity.Id
	}
	if binding := itemslot.GetBinding(); binding != nil {
		json["binding"] = binding.Key
	}

	return json
}

// func (i *Itemslot) Update(u *Universe, e *Entity, d time.Duration) {
// }
