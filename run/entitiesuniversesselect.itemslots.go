package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
	"ospokemon.com/space"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectItemslots)
}

func EntitiesUniversesSelectItemslots(args ...interface{}) {
	entities := args[0].(map[uint]*ospokemon.Entity)
	universe := args[1].(*ospokemon.Universe)
	itemslots, err := persistence.EntitiesItemsSelect(universe)

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
