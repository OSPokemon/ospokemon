package update

import (
	"github.com/ospokemon/ospokemon/world"
	"sort"
	"time"
)

func UpdateEntity(entity world.Entity, now time.Time) {
	entity.Controls().State = 0
	effects := world.Effects(make([]*world.Effect, 0))
	sort.Sort(entity.Effects())

	for _, effect := range entity.Effects() {
		UpdateEffect(effect, entity, now)

		if now.Before(effect.Start.Add(effect.Duration)) {
			effects = append(effects, effect)
		}
	}

	entity.SetEffects(effects)

	if entity.Controls().State&world.CTRLPnocast > 1 {
		entity.Controls().Action = nil
	} else if entity.Controls().Action != nil {
		UpdateAction(entity.Controls().Action, entity, now)
	}

	UpdateCollisions(entity)
}

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
