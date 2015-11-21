package world

import (
	// "math"
	"time"
)

type Entity interface {
	Name() string
	Physics() *Physics
	Graphics() *Graphics
	Controls() *Controls
	Effects() []*Effect
	SetEffects([]*Effect)
}

type Healthy interface {
	Health() int
	MaxHealth() int
	SetHealth(health int)
}

type Speedy interface {
	Speed() int
}

func UpdateEntity(id int, now time.Time) bool {
	entity := Entities[id]
	var update bool

	persistEffects := make([]*Effect, len(entity.Effects()))
	for _, effect := range entity.Effects() {
		update = effect.Update(id, entity, now) || update

		if !effect.Start.Add(effect.Duration).Before(now) {
			persistEffects = append(persistEffects, effect)
		}
	}
	entity.SetEffects(persistEffects)

	return update
}

func MoveEntity(id int, v *Vector) {
	entity := Entities[id]

	if int(entity.Controls().State&CTRLroot) > 0 {
		return
	}

	nextPos := entity.Physics().Position.Add(v)
	nextPhys := &Physics{
		Position: nextPos,
		Size:     entity.Physics().Size,
	}

	if int(entity.Controls().State&CTRLcollision) > 0 {
		for id2, entity2 := range Entities {
			if id == id2 {
				continue
			}

			if int(entity2.Controls().State&CTRLcollision) == 0 {
				continue
			}

			if nextPhys.CheckCollision(entity2.Physics()) {
				return
			}
		}
	}
}
