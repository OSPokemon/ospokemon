package world

import (
	"time"
)

type Entity interface {
	Name() string
	Physics() *Physics
	Graphics() *Graphics
	Controls() *Controls
	Effects() Effects
	SetEffects(effects Effects)
}

type Healthy interface {
	Health() int
	MaxHealth() int
	SetHealth(health int)
}

type Speedy interface {
	Speed() int
}

type Applicator interface {
	MakeEffects() Effects
}

type ArtificialIntelligence interface {
	Script() AiScript
}

type AiScript func(Entity, time.Time)
