// package data

// import (
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/ospokemon/ospokemon/objects/spellscripts"
// 	"github.com/ospokemon/ospokemon/world"
// 	"time"
// )

// type spellStore byte
// type abilitiesStore byte

// var SpellStore spellStore
// var AbilitiesStore abilitiesStore

// var Spells = make(map[int]*world.Spell)

// func (s *spellStore) Load(id int) *world.Spell {
// 	if Spells[id] != nil {
// 		return Spells[id]
// 	}

// 	row := Connection.QueryRow("SELECT id, name, casttime, cooldown, movecast, manacost, range, targettype, graphic FROM spells WHERE id=?", id)
// 	spell := &world.Spell{
// 		Cost: world.SpellCost{0, make(map[int]int)},
// 	}

// 	var casttime, cooldown int64
// 	err := row.Scan(&spell.Id, &spell.Name, &casttime, &cooldown, &spell.MoveCast, &spell.Cost.Mana, &spell.Range, &spell.TargetType, &spell.Graphic)
// 	spell.CastTime = time.Duration(casttime)
// 	spell.Cooldown = time.Duration(cooldown)
// 	spell.Script = spellscripts.Scripts[spell.Name]

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := Connection.Query("SELECT item_id, quantity FROM spell_reagents WHERE spell_id=?", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var item_id, quantity int
// 	for rows.Next() {
// 		err = rows.Scan(&item_id, &quantity)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		spell.Cost.Items[item_id] = quantity
// 	}

// 	Spells[id] = spell
// 	return spell
// }

// func (c abilitiesStore) GetForPokemon(id int) map[string]*world.Ability {
// 	rows, err := Connection.Query("SELECT spell_id, keybinding FROM pokemon_spells WHERE pokemon_id=?", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	abilities := make(map[string]*world.Ability)

// 	var spell_id int
// 	var keybinding string
// 	for rows.Next() {
// 		err = rows.Scan(&spell_id, &keybinding)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		abilities[keybinding] = &world.Ability{
// 			Spell: SpellStore.Load(spell_id),
// 		}
// 	}

// 	log.WithFields(log.Fields{
// 		"PokemonID": id,
// 		"Abilities": abilities,
// 	}).Debug("Controls loaded")

// 	return abilities
// }
