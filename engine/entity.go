package engine

import (
	"github.com/ospokemon/ospokemon/physics"
	"time"
)

type AiScript func(*Map, Entity, time.Time)

type Entity interface {
	EntityId() *int
	Graphic() *string
	Collision() *Collision
	Shape() physics.Shape
	SetShape(physics.Shape)
}

type livingEntity interface {
	Name() string
	Action() *Action
	SetAction(a *Action)
	Control() *Control
	Abilities() *[]*Ability
	Stats() map[string]*Stat
	Graphics() map[AnimationType]string
	Effects() *[]*Effect
	Walking() *physics.Point
	SetWalking(*physics.Point)
}

type aiEntity interface {
	AiScript() AiScript
}

type applicatorEntity interface {
	Apply(e Entity)
}

type LivingEntity interface {
	livingEntity
	Entity
}

type AiEntity interface {
	aiEntity
	Entity
}

type ApplicatorEntity interface {
	applicatorEntity
	Entity
}
