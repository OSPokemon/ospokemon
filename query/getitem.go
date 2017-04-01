package query

import (
	"ospokemon.com/game"
)

func GetItem(id uint) (*game.Item, error) {
	if spell, ok := game.Items[id]; ok {
		return spell, nil
	}

	return ItemsSelect(id)
}
