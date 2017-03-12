package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectTerrain)
}

func EntitiesUniversesSelectTerrain(args ...interface{}) {
	entity := args[0].(*game.Entity)
	universe := args[1].(*game.Universe)
	terrain, err := query.EntitiesTerrainSelect(entity, universe)

	if err == nil {
		entity.AddPart(terrain)

		imaging := game.MakeImaging()
		imaging.Image = terrain.Image
		entity.AddPart(imaging)
	} else if err.Error() != "sql: no rows in result set" {
		log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Error", err.Error()).Error("entity build terrain")
	}
}
