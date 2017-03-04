package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ActionCast, ActionCast)
}

func ActionCast(args ...interface{}) {
	u := args[0].(*game.Universe)
	e := args[1].(*game.Entity)
	a := args[2].(*game.Action)

	spell, err := query.GetSpell(a.Spell)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Universe": u.Id,
			"Entity":   e.Id,
			"Action":   a,
			"Error":    err.Error(),
		}).Error("action cast")
		return
	}

	timer := spell.Cooldown
	a.Timer = &timer

	if script, ok := game.Scripts[spell.Script]; ok {
		script(e, spell.Data)
	} else {
		logrus.WithFields(logrus.Fields{
			"Universe": u.Id,
			"Entity":   e.Id,
			"Action":   a,
			"Error":    err.Error(),
		}).Error("action cast")
	}
}
