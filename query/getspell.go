package query

import (
	"ospokemon.com/game"
)

func GetSpell(id uint) (*game.Spell, error) {
	if spell, ok := game.Spells[id]; ok {
		return spell, nil
	}

	return SpellsSelect(id)
}
