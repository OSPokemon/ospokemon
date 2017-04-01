package query

import (
	"ospokemon.com"
)

func GetItem(id uint) (*ospokemon.Item, error) {
	if spell, ok := ospokemon.Items[id]; ok {
		return spell, nil
	}

	return ItemsSelect(id)
}
