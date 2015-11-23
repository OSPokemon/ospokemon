package data

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
)

type spellStore byte
type abilityStore byte

var SpellStore spellStore
var AbilityStore abilityStore

var Spells = make(map[int]world.Spell)

func (s *spellStore) Load(id int) world.Spell {
	if Spells[id] != nil {
		return Spells[id]
	}

	row := Connection.QueryRow("SELECT id, name, t, category, description, pp, power, accuracy, targeter, casttime, cooldown, contestcategory, appeal, jam, contestdescription, priority FROM moves WHERE id=?", id)
	spell := &MoveSpell{
		COST: &world.SpellCost{},
	}

	err := row.Scan(&spell.ID, &spell.NAME, &spell.T, &spell.CATEGORY, &spell.DESCRIPTION, &spell.Pp, &spell.POWER, &spell.ACCURACY, &spell.TARGETER, &spell.CASTTIME, &spell.COOLDOWN, &spell.CONTESTCATEGORY, &spell.APPEAL, &spell.JAM, &spell.CONTESTDESCRIPTION, &spell.PRIORITY)
	spell.COST.Mana = spell.PP()
	// TODO spell reagent costs

	if err != nil {
		log.Fatal(err)
	}

	Spells[id] = spell
	return spell
}

func (s *spellStore) FetchIdsForPokemon(id int) []int {
	rows, err := Connection.Query("SELECT move_id FROM moves_pokemon WHERE pokemon_id=?", id)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	spell_ids := make([]int, 0)
	var spell_id int

	for rows.Next() {
		err = rows.Scan(&spell_id)

		if err != nil {
			log.Fatal(err)
		}

		spell_ids = append(spell_ids, spell_id)
	}

	return spell_ids
}

func (a *abilityStore) New(id int) *world.Ability {
	spell := SpellStore.Load(id)
	return &world.Ability{
		Spell: spell,
	}
}
