package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"time"
)

func SpellsSelect(id uint) (*game.Spell, error) {
	row := Connection.QueryRow(
		"SELECT script, casttime, cooldown FROM spells WHERE id=?",
		id,
	)

	spell := game.MakeSpell(id)

	var casttimebuff, cooldownbuff int64
	if err := row.Scan(&spell.Script, &casttimebuff, &cooldownbuff); err != nil {
		return spell, err
	}

	if t := time.Duration(casttimebuff); casttimebuff > 0 {
		spell.CastTime = t * time.Millisecond
	}
	if t := time.Duration(cooldownbuff); cooldownbuff > 0 {
		spell.Cooldown = t * time.Millisecond
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_spells WHERE spell=?",
		id,
	)
	if err != nil {
		return spell, err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return spell, err
		}
		spell.Animations[keybuff] = valuebuff
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT key, value FROM spells_data WHERE spell=?",
		id,
	)
	if err != nil {
		return spell, err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return spell, err
		}
		spell.Data[keybuff] = valuebuff
	}
	rows.Close()

	game.Spells[id] = spell

	logrus.WithFields(logrus.Fields{
		"Spell": spell.Id,
	}).Info("spells select")

	return spell, nil
}
