package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.UniversesSelect, UniversesSelectEntities)
}

func UniversesSelectEntities(args ...interface{}) {
	universe := args[0].(*game.Universe)
	entities, err := query.EntitiesUniversesSelect(universe)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Universe": universe.Id,
			"Error":    err.Error(),
		}).Error("universes select entities")
		return
	}

	for _, entity := range entities {
		universe.Add(entity)
	}
}
