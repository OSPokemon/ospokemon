package update

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
	"sort"
	"time"
)

func UpdateEntity(entity world.Entity, now time.Time) {

	if applicator, ok := entity.(world.Applicator); ok {
		updateApplicatorEntity(applicator, now)
	}

	if mortal, ok := entity.(world.Mortality); ok {
		updateMortalEntity(mortal, now)
	}

	if intelligent, ok := entity.(world.Intelligence); ok {
		updateIntelligentEntity(intelligent, now)
	}
}

func updateApplicatorEntity(entity world.Applicator, now time.Time) {
	for _, entity2 := range world.Entities {
		if entity == entity2 {
			continue
		}
		if !entity2.Physics().Solid {
			continue
		}

		entity.Apply(entity2)
	}
}

func updateMortalEntity(entity world.Mortality, now time.Time) {
	resetMortalEntity(entity, now)

	applyEffectScripts(entity, now)
}

func updateIntelligentEntity(entity world.Intelligence, now time.Time) {
	resetIntelligentEntity(entity, now)

	entity.Script()

	moved := maybeWalk(entity, now)

	maybeCast(entity, now, moved)
}

func resetMortalEntity(entity world.Mortality, now time.Time) {
	entity.SetControl(entity.Control() & world.CTRLdead) // dead flag doesn't reset. dead is not a debuff

	persisteffects := make([]*world.Effect, 0)
	for _, effect := range entity.Effects() {
		if effect.Start == nil {
			log.WithFields(log.Fields{
				"entity": entity.Name(),
				"effect": effect,
			}).Debug("Effect timer starts")
			effect.Start = &now
		}
		if effect.Start.Add(effect.Duration).Before(now) {
			log.WithFields(log.Fields{
				"entity": entity.Name(),
				"effect": effect,
			}).Debug("Effect falls off")
		} else {
			persisteffects = append(persisteffects, effect)
		}
	}
	entity.SetEffects(persisteffects)

	for _, stat := range entity.Stats() {
		if stat.Value() > stat.MaxValue() {
			stat.SetValue(stat.MaxValue())
		}
		stat.SetMaxValue(stat.BaseMaxValue())
	}
}

func applyEffectScripts(entity world.Mortality, now time.Time) {
	effects := entity.Effects()
	sort.Sort(fxsorter(effects))

	for _, effect := range effects {
		log.WithFields(log.Fields{
			"entity": entity.Name(),
			"effect": effect,
		}).Debug("Effect tick")
		effect.Script(effect, entity, now)
	}
}

func resetIntelligentEntity(entity world.Intelligence, now time.Time) {
	for _, ability := range entity.Abilities() {
		ability.CastTime = ability.Spell.CastTime
		ability.Cooldown = ability.Spell.Cooldown
		ability.MoveCast = ability.Spell.MoveCast
		ability.Cost = ability.Spell.Cost
		ability.Range = ability.Spell.Range
		ability.TargetData = tdcopier(ability.Spell.TargetData).copy()
	}
}

func maybeWalk(entity world.Intelligence, now time.Time) bool {
	if world.IsDead(entity) {
		return false // dead guys dont walk
	}
	if world.IsStuck(entity) {
		return false
	}

	destination := entity.Walking()

	if destination == nil {
		return false
	}
	if entity.Action() != nil && entity.Action().Start == nil && !entity.Action().Ability.MoveCast {
		entity.SetWalking(nil)
		return false
	}

	speed := entity.Stats()["speed"].Value()
	distance := world.GetDistance(&entity.Physics().Point, destination)

	if float64(speed) > distance {
		speed = int(distance)
		entity.SetWalking(nil)
	}

	vector := world.CreatePathVector(&entity.Physics().Point, destination, speed)
	entity.Graphics().Current = entity.Graphics().Animations[vector.AnimationType()]

	MoveEntity(entity, vector)
	return true
}

func maybeCast(entity world.Intelligence, now time.Time, moved bool) {
	if world.IsDead(entity) {
		return // dead guys dont cast
	}
	if world.NoCast(entity) {
		return
	}
	if entity.Action() == nil {
		return
	}
	if !entity.Action().Ability.MoveCast && moved {
		entity.SetAction(nil)
		return
	}

	if entity.Action().Start == nil {
		if entity.Action().Ability.LastCast.Add(entity.Action().Ability.Cooldown).After(now) {
			return
		}

		entity.Action().Start = &now
	}
	if entity.Action().Start.Add(entity.Action().Ability.CastTime).Before(now) {
		log.WithFields(log.Fields{
			"entity":  entity.Name(),
			"ability": entity.Action().Ability.Spell.Name,
		}).Debug("Cast time complete")

		if entity.Action().Ability.Spell.Script == nil {
			log.Warn("ABORT CAST")
		}

		entity.Action().Ability.Spell.Script(entity, entity.Action().Target, now)
		entity.Action().Ability.LastCast = now
		entity.SetAction(nil)
	}
}

// target data can be coppied

type tdcopier map[string]interface{}

func (src tdcopier) copy() map[string]interface{} {
	dst := make(map[string]interface{})

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

// Effects can be sorted

type fxsorter []*world.Effect

func (e fxsorter) Len() int {
	return len(e)
}

func (e fxsorter) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e fxsorter) Less(i, j int) bool {
	return e[i].Priority < e[j].Priority
}
