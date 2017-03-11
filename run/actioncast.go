package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func init() {
	event.On(event.ActionCast, ActionCast)
}

func ActionCast(args ...interface{}) {
	universe := args[0].(*game.Universe)
	entity := args[1].(*game.Entity)
	action := args[2].(*game.Action)

	spell := action.Spell
	timer := spell.Cooldown
	action.Timer = &timer

	if script, ok := game.Scripts[spell.Script]; ok {
		if err := script(entity, spell.Data); err != nil {
			logrus.WithFields(logrus.Fields{
				"Universe": universe.Id,
				"Entity":   entity.Id,
				"Action":   action,
				"Error":    err.Error(),
			}).Error("action cast")
		}
	}
}
