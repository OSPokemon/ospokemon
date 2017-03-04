package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetItem(id uint) (*game.Item, error) {
	if spell, ok := game.Items[id]; ok {
		return spell, nil
	}

	return ItemsSelect(id)
}
