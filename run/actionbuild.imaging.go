package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.ActionBuild, ActionBuildImaging)
}

func ActionBuildImaging(args ...interface{}) {
	action := args[0].(*game.Action)

	spell, err := query.GetSpell(action.Spell)

	if spell != nil {
		imaging := game.MakeImaging()
		imaging.ReadAnimations(spell.Animations)
		action.AddPart(imaging)
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"Spell": action.Spell,
			"Error": err.Error(),
		}).Error("action build imaging")
	} else {
		logrus.WithFields(logrus.Fields{
			"Spell": action.Spell,
		}).Warn("action build imaging")
	}
}
