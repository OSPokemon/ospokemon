package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectSpawners)
}

func EntitiesUniversesSelectSpawners(args ...interface{}) {
	entities := args[0].(map[uint]*game.Entity)
	universe := args[1].(*game.Universe)
	spawners, err := query.EntitiesSpawnersSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("entities universes select spawners")
		return
	}

	for entityId, spawner := range spawners {
		spawner.Child = entities[entityId]
		spawner.Child.AddPart(spawner)
		universe.AddSpawner(spawner)
	}
}
