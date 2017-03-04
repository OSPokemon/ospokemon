package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetSpell(id uint) (*game.Spell, error) {
	if spell, ok := game.Spells[id]; ok {
		return spell, nil
	}

	return SpellsSelect(id)
}
