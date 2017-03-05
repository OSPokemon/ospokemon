package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectItemslot)
}

func EntitiesUniversesSelectItemslot(args ...interface{}) {
	entity := args[0].(*game.Entity)
	universe := args[1].(*game.Universe)
	itemslot, err := query.EntitiesItemsSelect(entity, universe)

	if err == nil {
		entity.AddPart(itemslot)
		entity.AddPart(itemslot.Parts[part.Imaging])
		itemslot.Parts = entity.Parts
	} else if err.Error() != "sql: no rows in result set" {
		logrus.WithFields(logrus.Fields{
			"Universe": universe.Id,
			"Entity":   entity.Id,
			"Error":    err.Error(),
		}).Error("run.EntitiesUniversesSelectItemslot")
	}
}
