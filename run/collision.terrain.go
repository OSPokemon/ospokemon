package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.Collision, func(args ...interface{}) {
		entity1 := args[0].(*game.Entity)
		entity2 := args[1].(*game.Entity)
		universe := args[2].(*game.Universe)
		vector := args[3].(space.Vector)

		terrain := entity2.GetTerrain()
		if terrain == nil {
			return
		}

		if terrain.Collision {
			vector = space.Line{
				entity2.Shape.Center(),
				entity1.Shape.Center(),
			}.Vector().MakeUnit().Multiply(vector.Length())
			entity1.Move(vector, universe)
		}
	})
}
