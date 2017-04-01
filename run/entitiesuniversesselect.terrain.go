package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectTerrain)
}

func EntitiesUniversesSelectTerrain(args ...interface{}) {
	entities := args[0].(map[uint]*ospokemon.Entity)
	universe := args[1].(*ospokemon.Universe)
	terrains, err := query.EntitiesTerrainsSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("entities universes select terrain")
	}

	for entityId, terrain := range terrains {
		entity := entities[entityId]
		entity.AddPart(terrain)

		imaging := ospokemon.MakeImaging()
		imaging.Image = terrain.Image
		entity.AddPart(imaging)
	}
}
