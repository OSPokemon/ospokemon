package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

type spellStore byte
type controlsStore byte

var SpellStore spellStore
var ControlsStore controlsStore

var Spells = make(map[int]*world.Spell)

func (s *spellStore) Load(id int) *world.Spell {
	if Spells[id] != nil {
		return Spells[id]
	}

	row := Connection.QueryRow("SELECT id, name, casttime, cooldown, movecast, manacost, range, targettype FROM spells WHERE id=?", id)
	spell := &world.Spell{
		Cost: world.SpellCost{0, make(map[int]int)},
	}

	err := row.Scan(&spell.Id, &spell.Name, &spell.CastTime, &spell.Cooldown, &spell.MoveCast, &spell.Cost.Mana, &spell.Range, &spell.TargetType)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := Connection.Query("SELECT item_id, quantity FROM spell_reagents WHERE spell_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var item_id, quantity int
	for rows.Next() {
		err = rows.Scan(&item_id, &quantity)
		if err != nil {
			log.Fatal(err)
		}

		spell.Cost.Items[item_id] = quantity
	}

	Spells[id] = spell
	return spell
}

func (c controlsStore) BuildForPokemon(id int) *world.Controls {
	rows, err := Connection.Query("SELECT spell_id, keybinding FROM pokemon_spells WHERE pokemon_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	controls := &world.Controls{nil, 0, make(map[string]*world.Ability)}

	var spell_id int
	var keybinding string
	for rows.Next() {
		err = rows.Scan(&spell_id, &keybinding)
		if err != nil {
			log.Fatal(err)
		}

		controls.Abilities[keybinding] = &world.Ability{
			Spell: SpellStore.Load(spell_id),
		}
	}

	return controls
}
