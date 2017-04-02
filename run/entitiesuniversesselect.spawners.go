package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectSpawners)
}

func EntitiesUniversesSelectSpawners(args ...interface{}) {
	entities := args[0].(map[uint]*ospokemon.Entity)
	universe := args[1].(*ospokemon.Universe)
	spawners, err := persistence.EntitiesSpawnersSelect(universe)

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
