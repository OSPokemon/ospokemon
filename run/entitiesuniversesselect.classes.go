package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectClasses)
}

func EntitiesUniversesSelectClasses(args ...interface{}) {
	entities := args[0].(map[uint]*game.Entity)
	universe := args[1].(*game.Universe)
	classes, err := query.ClassesEntitiesSelect(universe)

	if err != nil {
		log.Add("Universe", universe.Id).Add("Error", err.Error()).Error("entities universes select classes")
		return
	}

	for entityId, class := range classes {
		entity := entities[entityId]

		entity.AddPart(game.BuildImaging(class.Animations))

		rect := entity.Shape.(*space.Rect)
		rect.Dimension.DX = class.Dimension.DX
		rect.Dimension.DY = class.Dimension.DY
	}
}
