package json

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func Itemslot(i *game.Itemslot) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"id":     i.Id,
		"amount": i.Amount,
	}

	if item, _ := query.GetItem(i.Item); item != nil {
		data["item"] = item.Snapshot()
	} else {
		data["item"] = i.Item
	}

	for _, part := range i.Parts {
		if imaging, ok := part.(*game.Imaging); ok {
			key, partData := Imaging(imaging)
			data[key] = partData
		}
	}

	return "itemslot", data
}
