package query

import (
	"ospokemon.com"
)

func GetSpell(id uint) (*ospokemon.Spell, error) {
	if spell, ok := ospokemon.Spells[id]; ok {
		return spell, nil
	}

	return SpellsSelect(id)
}
