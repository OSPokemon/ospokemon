package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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
		entity.AddPart(itemslot.GetImaging())
		itemslot.Parts = entity.Parts
	} else if err.Error() != "sql: no rows in result set" {
		log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Error", err.Error()).Error("run.EntitiesUniversesSelectItemslot")
	}
}
