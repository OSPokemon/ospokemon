package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

var Spells = make(map[int]*world.Spell)

func LoadSpell(id int) {
	if Spells[id] != nil {
		return
	}

	row := Connection.QueryRow("SELECT id, name, casttime, cooldown, movecast, manacost, range, targettype, graphic FROM spells WHERE id=?", id)
	importSpellRecord(row)

	rows, err := Connection.Query("SELECT spell_id, item_id, quantity FROM spell_reagents WHERE spell_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		importSpellReagentRecord(rows)
	}

	log.WithFields(log.Fields{
		"id":    id,
		"spell": Spells[id],
	}).Debug("Spell built")
}

func LoadAllSpells() {
	rows, err := Connection.Query("SELECT id, name, casttime, cooldown, movecast, manacost, range, targettype, graphic FROM spells")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		importSpellRecord(rows)
	}

	rows, err = Connection.Query("SELECT spell_id, item_id, quantity FROM spell_reagents")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		importSpellReagentRecord(rows)
	}

	log.Debug("All spells built")
	log.Debug(Spells)
}

func MakeSpellsForPlayer(id int) map[string]*world.Spell {
	spells := make(map[string]*world.Spell)

	rows, err := Connection.Query("SELECT spell_base_id, name, keybinding FROM players_spells WHERE player_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var spell_base_id int
	var name, keybinding string
	for rows.Next() {
		err = rows.Scan(&spell_base_id, &name, &keybinding)
		if err != nil {
			log.Fatal(err)
		}

		spell := Spells[spell_base_id]

		if spell == nil {
			log.WithFields(log.Fields{
				"PlayerId": id,
				"SpellId":  spell_base_id,
			}).Warn("Invalid reference for spell")
			continue
		}
		if spell.Script == nil {
			log.WithFields(log.Fields{
				"PlayerId": id,
				"SpellId":  spell_base_id,
			}).Warn("Invalid reference for spell script")
			continue
		}

		spell = makeSpellCopy(spell)

		spell.Name = name

		rows, err = Connection.Query("SELECT key, value FROM players_spells_targetdata WHERE player_id=? AND spell_base_id=? AND keybinding=?", id, spell_base_id, keybinding)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var key, value string
		for rows.Next() {
			rows.Scan(&key, &value)
			spell.TargetData[key] = value
		}

		spells[keybinding] = spell
	}

	return spells
}

func importSpellRecord(row Scannable) {
	spell := &world.Spell{
		Cost: world.SpellCost{0, make(map[int]int)},
	}
	var casttime, cooldown int64

	err := row.Scan(&spell.Id, &spell.Name, &casttime, &cooldown, &spell.MoveCast, &spell.Cost.Mana, &spell.Range, &spell.TargetType, &spell.Graphic)
	if err != nil {
		log.Fatal(err)
	}

	spell.CastTime = time.Duration(casttime)
	spell.Cooldown = time.Duration(cooldown)
	spell.Script = registry.Scripts[spell.Name]

	if spell.Script == nil {
		log.WithFields(log.Fields{
			"SpellId": spell.Id,
		}).Warn("Spell loaded with missing script")
	}

	Spells[spell.Id] = spell
}

func importSpellReagentRecord(row Scannable) {
	var spell_id, item_id, quantity int
	err := row.Scan(&spell_id, &item_id, &quantity)
	if err != nil {
		log.Fatal(err)
	}

	Spells[spell_id].Cost.Items[item_id] = quantity
}

func makeSpellCopy(s *world.Spell) *world.Spell {
	return &world.Spell{
		Id:         s.Id,
		Name:       s.Name,
		CastTime:   s.CastTime,
		Cooldown:   s.Cooldown,
		MoveCast:   s.MoveCast,
		Cost:       s.Cost,
		Range:      s.Range,
		TargetType: s.TargetType,
		TargetData: tdcopier(s.TargetData).copy(),
		Graphic:    s.Graphic,
		Script:     s.Script,
	}
}

// target data can be coppied

type tdcopier map[string]interface{}

func (src tdcopier) copy() map[string]interface{} {
	dst := make(map[string]interface{})

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
