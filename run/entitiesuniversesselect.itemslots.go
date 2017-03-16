package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectItemslots)
}

func EntitiesUniversesSelectItemslots(args ...interface{}) {
	entities := args[0].(map[uint]*game.Entity)
	universe := args[1].(*game.Universe)
	itemslots, err := query.EntitiesItemsSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("run.EntitiesUniversesSelectItemslots")
		return
	}

	for entityId, itemslot := range itemslots {
		entity := entities[entityId]
		entity.AddPart(itemslot)
		entity.AddPart(itemslot.GetImaging())
		itemslot.Parts = entity.Parts

		rect := entity.Shape.(*space.Rect)
		item := itemslot.Item
		rect.Dimension.DX = item.Dimension.DX
		rect.Dimension.DY = item.Dimension.DY
	}
}
