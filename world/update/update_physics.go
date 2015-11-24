package update

import (
	"github.com/ospokemon/ospokemon/world"
)

func MoveEntity(entity world.Entity, v *world.Vector) {
	if entity.Controls().State&world.CTRLPstuck > 0 {
		return
	}

	nextPos := entity.Physics().Position.Add(v)

	if entity.Physics().Solid {
		nextPhys := &world.Physics{
			Position: nextPos,
			Size:     entity.Physics().Size,
			Solid:    entity.Physics().Solid,
		}

		for _, entity2 := range world.Entities {
			if entity == entity2 {
				continue
			}
			if !entity2.Physics().Solid {
				continue
			}

			if nextPhys.CheckCollision(entity2.Physics()) {
				return
			}
		}
	}

	entity.Physics().Position = nextPos
}

func UpdateCollisions(entity world.Entity) {
	for _, entity2 := range world.Entities {
		if entity == entity2 {
			continue
		}
		if entity2.Physics().Solid {
			continue
		}

		if applicator, ok := entity2.(world.Applicator); ok {
			for _, effect := range applicator.MakeEffects() {
				entity.SetEffects(append(entity.Effects(), effect))
			}
		}
	}
}
