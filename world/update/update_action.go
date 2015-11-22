package update

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

func UpdateAction(action *world.Action, entity world.Entity, now time.Time) {
	if action.Ability == world.WalkAbility {
		position := entity.Physics().Position
		destination := action.Target.(*world.Position)

		if world.GetDistance(&position, destination) < 10 {
			entity.Controls().Action = nil
			return
		}

		speedy := entity.(world.Speedy)
		vector := world.CreatePathVector(&entity.Physics().Position, destination, speedy.Speed())

		moveEffect := &world.Effect{"walk", world.EFCTmove, vector, now, 0}
		entity.SetEffects(append(entity.Effects(), moveEffect))

		log.Printf("MoveAction %s along %v towards %v", entity.Name(), vector, destination)
		return
	}
}
