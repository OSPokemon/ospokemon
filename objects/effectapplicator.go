package objects

import (
	"github.com/ospokemon/ospokemon/aiscripts"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/physics"
	"time"
)

type EffectApplicator struct {
	ENTITYID   int
	GRAPHIC    string
	COLLISION  engine.Collision
	SHAPE      physics.Shape
	SCRIPT     engine.AiScript
	APPLY      func(engine.Entity)
	EXPIRATION time.Time
}

func NewEffectApplicator(shape physics.Shape, apply func(engine.Entity), graphic string, expiration time.Time) *EffectApplicator {
	return &EffectApplicator{
		GRAPHIC:    graphic,
		COLLISION:  engine.CLSNfluid,
		SHAPE:      shape,
		SCRIPT:     aiscripts.MaybeExpireEntity,
		APPLY:      apply,
		EXPIRATION: expiration,
	}
}

func (e *EffectApplicator) EntityId() *int {
	return &e.ENTITYID
}

func (e *EffectApplicator) Graphic() *string {
	return &e.GRAPHIC
}

func (e *EffectApplicator) Collision() *engine.Collision {
	return &e.COLLISION
}

func (e *EffectApplicator) Shape() physics.Shape {
	return e.SHAPE
}

func (e *EffectApplicator) SetShape(shape physics.Shape) {
	e.SHAPE = shape
}

func (e *EffectApplicator) Apply(entity engine.Entity) {
	e.APPLY(entity)
}

func (e *EffectApplicator) Expire() time.Time {
	return e.EXPIRATION
}
