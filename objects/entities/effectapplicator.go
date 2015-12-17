package entities

import (
	"github.com/ospokemon/ospokemon/objects/aiscripts"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type EffectApplicator struct {
	ENTITYID   int
	NAME       string
	PHYSICS    *world.Physics
	GRAPHICS   *world.Graphics
	SCRIPT     world.AiScript
	APPLY      func(world.Entity)
	EXPIRATION time.Time
}

var effectapplicatorcontrolsstub uint8
var effectapplicatorstatsstub = make(map[string]world.Stat)
var effectapplicatoreffectsstub = make([]*world.Effect, 0)
var effectapplicatorabilitiesstub = make(map[string]*world.Ability)

func NewEffectApplicator(name string, shape physics.Shape, apply func(world.Entity), graphic string, expiration time.Time) *EffectApplicator {
	return &EffectApplicator{
		NAME: name,
		PHYSICS: &world.Physics{
			Shape: shape,
			Solid: false,
		},
		GRAPHICS: &world.Graphics{
			Current:    graphic,
			Animations: make(map[world.AnimationType]string),
		},
		SCRIPT:     aiscripts.MaybeExpireEntity,
		APPLY:      apply,
		EXPIRATION: expiration,
	}
}

func (e *EffectApplicator) Apply(ent world.Entity) {
	e.APPLY(ent)
}

func (e *EffectApplicator) EntityId() int {
	return e.ENTITYID
}

func (e *EffectApplicator) SetEntityId(id int) {
	e.ENTITYID = id
}

func (e *EffectApplicator) Name() string {
	return e.NAME
}

func (e *EffectApplicator) Physics() *world.Physics {
	return e.PHYSICS
}

func (e *EffectApplicator) Graphics() *world.Graphics {
	return e.GRAPHICS
}

func (e *EffectApplicator) Action() *world.Action {
	return nil
}

func (e *EffectApplicator) SetAction(a *world.Action) {}

func (e *EffectApplicator) Control() uint8 {
	return effectapplicatorcontrolsstub
}

func (e *EffectApplicator) SetControl(uint8) {}

func (e *EffectApplicator) Stats() map[string]world.Stat {
	return effectapplicatorstatsstub
}

func (e *EffectApplicator) Effects() []*world.Effect {
	return effectapplicatoreffectsstub
}

func (e *EffectApplicator) SetEffects([]*world.Effect) {}

func (e *EffectApplicator) Script() world.AiScript {
	return aiscripts.MaybeExpireEntity
}

func (e *EffectApplicator) Walking() *physics.Point {
	return nil
}

func (e *EffectApplicator) SetWalking(p *physics.Point) {}

func (e *EffectApplicator) Abilities() map[string]*world.Ability {
	return effectapplicatorabilitiesstub
}

func (e *EffectApplicator) Expiration() time.Time {
	return e.EXPIRATION
}
