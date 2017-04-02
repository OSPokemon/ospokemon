package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
	"ospokemon.com/space"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectClasses)
}

func EntitiesUniversesSelectClasses(args ...interface{}) {
	entities := args[0].(map[uint]*ospokemon.Entity)
	universe := args[1].(*ospokemon.Universe)
	classes, err := persistence.ClassesEntitiesSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("entities universes select classes")
		return
	}

	for entityId, class := range classes {
		entity := entities[entityId]

		entity.AddPart(ospokemon.BuildImaging(class.Animations))

		rect := entity.Shape.(*space.Rect)
		rect.Dimension.DX = class.Dimension.DX
		rect.Dimension.DY = class.Dimension.DY
	}
}
