package engine

import (
	"time"
)

func UpdateEntity(m *Map, e Entity, now time.Time, t time.Duration) {
	if applicator, ok := e.(ApplicatorEntity); ok {
		UpdateApplicatorEntity(m, applicator, now, t)
	}

	if living, ok := e.(LivingEntity); ok {
		UpdateLivingEntity(m, living, now, t)
	}

	if ai, ok := e.(AiEntity); ok {
		UpdateAiEntity(m, ai, now, t)
	}
}

func UpdateApplicatorEntity(m *Map, applicator ApplicatorEntity, now time.Time, t time.Duration) {
	for _, e := range m.Entities {
		if e == applicator {
			continue
		}
		if *e.Collision() == CLSNnone {
			continue
		}

		if applicator.Shape() == nil || e.Shape() == nil || CheckCollision(applicator.Shape(), e.Shape()) {
			applicator.Apply(e)
		}
	}
}

func UpdateLivingEntity(m *Map, e LivingEntity, now time.Time, t time.Duration) {
	for _, stat := range e.Stats() {
		stat.Max = stat.Base
		stat.Regen = stat.RegenBase
	}

	for _, ability := range *e.Abilities() {
		ability.CastTime = ability.Spell.CastTime
		ability.MoveCast = ability.Spell.MoveCast
		ability.ChannelTime = ability.Spell.ChannelTime
		ability.MoveChannel = ability.Spell.MoveChannel
		ability.Cooldown = ability.Spell.Cooldown
		ability.ManaCost = ability.Spell.ManaCost
		ability.ItemCost = ability.Spell.ItemCost
		ability.TargetType = ability.Spell.TargetType
		ability.Range = ability.Spell.Range
		ability.Size = ability.Spell.Size
	}

	effects := make([]*Effect, 0)
	for _, effect := range *e.Effects() {
		if effect.Start == nil {
			effect.Start = &now
		}

		effect.EffectScript(effect, m, e, now)

		if effect.Start.Add(effect.Duration).After(now) {
			effects = append(effects, effect)
		}
	}
	*e.Effects() = effects

	for _, stat := range e.Stats() {
		if stat.Value > stat.Max {
			stat.Value = stat.Max
		}
	}

	if walking := e.Walking(); walking != nil && *e.Control()&CTRLPstuck < 1 {
		if action := e.Action(); action != nil && action.CastStart == nil && !action.Ability.MoveCast {
			e.SetWalking(nil)
		} else {
			speed := e.Stats()["speed"].Value
			distance := distancePointShape(*walking, e.Shape())

			if speed > distance {
				speed = distance
				e.SetWalking(nil)
			}

			vector := CreatePathVector(e.Shape(), *walking, int(speed))
			*e.Graphic() = e.Graphics()[DirectionAnimation(vector)]
			MoveEntity(m, e, vector)
		}
	}

	if action := e.Action(); action != nil && *e.Control()&CTRLPnocast < 1 && (e.Walking() == nil || action.Ability.MoveCast) {
		if action.CastStart == nil {
			action.CastStart = &now
		}

		if !action.CastStart.Add(action.Ability.CastTime).After(now) && (e.Walking() == nil || action.Ability.MoveChannel) {
			if action.ChannelStart == nil {
				action.ChannelStart = &now
			}

			action.Ability.Spell.Script(m, e, action.Target, now)

			if !action.ChannelStart.Add(action.Ability.ChannelTime).After(now) {
				e.SetAction(nil)
			}
		}
	}
}

func UpdateAiEntity(m *Map, e AiEntity, now time.Time, t time.Duration) {
	e.AiScript()(m, e, now)
}
