package update

import (
	"github.com/ospokemon/ospokemon/world"
	"sort"
	"time"
)

func ResetEntity(entity world.Entity) {
	entity.Controls().State = 0

	for _, ability := range entity.Controls().Abilities {
		ability.CastTime = ability.Spell.CastTime
		ability.Cooldown = ability.Spell.Cooldown
		ability.MoveCast = ability.Spell.MoveCast
		ability.Cost = ability.Spell.Cost
		ability.Range = ability.Spell.Range
	}
}

func UpdateEntity(entity world.Entity, now time.Time) {
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
	}

	if entity.Controls().Action != nil {
		if !entity.Controls().Action.Ability.MoveCast {
			entity.Physics().Walking = nil
		}

		entity.Controls().Action.Ability.Spell.Script(entity, entity.Controls().Action.Target, now)
	}

	if speedy, ok := entity.(world.Speedy); ok && entity.Physics().Walking != nil {
		speed := speedy.Speed()
		destination := entity.Physics().Walking
		distance := world.GetDistance(&entity.Physics().Position, destination)

		if float64(speed) > distance {
			speed = int(distance)
			entity.Physics().Walking = nil
		}

		vector := world.CreatePathVector(&entity.Physics().Position, destination, speed)
		MoveEntity(entity, vector)
	}
}
