package world

import (
	"github.com/ospokemon/ospokemon/physics"
	"time"
)

type AiScript func(Entity, time.Time)

type Entity interface {
	EntityId() int
	SetEntityId(id int)
	Name() string
	Physics() *Physics
	Graphics() *Graphics
	Action() *Action
	SetAction(a *Action)
}

type Applicator interface {
	Entity
	Apply(Entity)
}

type mortality interface {
	Control() uint8
	SetControl(uint8)
	Stats() map[string]Stat
	Effects() []*Effect
	SetEffects([]*Effect)
}

type Mortality interface {
	Entity
	mortality
}

type Intelligence interface {
	Entity
	mortality
	Script() AiScript
	Walking() *physics.Point
	SetWalking(p *physics.Point)
	Abilities() map[string]*Ability
}

type Stat interface {
	Value() int
	SetValue(value int)
	MaxValue() int
	SetMaxValue(maxvalue int)
	BaseMaxValue() int
	SetBaseMaxValue(basemaxvalue int)
}
