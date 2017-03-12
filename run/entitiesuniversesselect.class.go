package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.EntitiesUniversesSelect, EntitiesUniversesSelectClass)
}

func EntitiesUniversesSelectClass(args ...interface{}) {
	entity := args[0].(*game.Entity)
	universe := args[1].(*game.Universe)
	class, err := query.ClassesEntitiesSelect(entity, universe)

	if err == nil {
		imaging := game.MakeImaging()
		imaging.ReadAnimations(class.Animations)
		imaging.Image = imaging.Animations["portrait"]

		entity.AddPart(imaging)

		rect := entity.Shape.(*space.Rect)
		rect.Dimension.DX = class.Dimension.DX
		rect.Dimension.DY = class.Dimension.DY
	} else if err.Error() != "sql: no rows in result set" {
		log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Error", err.Error()).Error("entity build class")
	}
}
