package util

import (
	"errors"
	"ospokemon.com"
	"strconv"
)

func GetSpell(data interface{}) (*ospokemon.Spell, error) {
	var spell *ospokemon.Spell
	var err error

	switch data_spell := data.(type) {
	case nil:
		err = errors.New("getspell: nil")
	case *ospokemon.Spell:
		spell = data_spell
	case int:
		spell, err = ospokemon.GetSpell(uint(data_spell))
	case uint:
		spell, err = ospokemon.GetSpell(data_spell)
	case float64:
		spell, err = ospokemon.GetSpell(uint(data_spell))
	case string:
		spellid64, err2 := strconv.ParseUint(data_spell, 10, 0)
		if err2 == nil {
			spell, err = ospokemon.GetSpell(uint(spellid64))
		} else {
			err = errors.New("getitem: " + err2.Error())
		}
	default:
		err = errors.New("getitem: \"spell\" type unrecognized")
	}

	return spell, err
}
