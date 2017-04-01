package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.UniversesSelect, UniversesSelectEntities)
}

func UniversesSelectEntities(args ...interface{}) {
	universe := args[0].(*ospokemon.Universe)
	entities, err := query.EntitiesUniversesSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("universes select entities")
		return
	}

	for _, entity := range entities {
		universe.Add(entity)
	}
}
