package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectDialog)
}

func EntitiesUniversesSelectDialog(args ...interface{}) {
	entity := args[0].(*game.Entity)
	universe := args[1].(*game.Universe)
	dialog, err := query.DialogsSelect(entity, universe)

	if err == nil {
		entity.AddPart(dialog)
	} else if err.Error() != "sql: no rows in result set" {
		logrus.WithFields(logrus.Fields{
			"Universe": universe.Id,
			"Entity":   entity.Id,
			"Error":    err.Error(),
		}).Error("EntityUniverseSelect-Dialog")
	}
}
