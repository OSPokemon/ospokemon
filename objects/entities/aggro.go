package entities

import (
	// "github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/data"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type AggressionLevel bool

const (
	AGGROpassive AggressionLevel = false
	AGGROeager                   = true
)

type AiProfile struct {
	AggressionLevel
	IdleRange    float64
	ChaseRange   float64
	HomePosition *world.Position
	Threat       map[int]int
	Target       int
}

type AiPokemonEntity struct {
	Entity  data.PokemonEntity
	Profile *AiProfile
}

// PokemonEntity is an entity

func (p *AiPokemonEntity) Name() string {
	return p.Entity.Name()
}

func (p *AiPokemonEntity) Physics() *world.Physics {
	return p.Entity.Physics()
}

func (p *AiPokemonEntity) Graphics() *world.Graphics {
	return p.Entity.Graphics()
}

func (p *AiPokemonEntity) Action() *world.Action {
	return p.Entity.Action()
}

func (p *AiPokemonEntity) SetAction(action *world.Action) {
	p.Entity.SetAction(action)
}

// AiPokemonEntity is mortal

func (p *AiPokemonEntity) Control() uint8 {
	return p.Entity.Control()
}

func (p *AiPokemonEntity) SetControl(control uint8) {
	p.Entity.SetControl(control)
}

func (p *AiPokemonEntity) Stats() map[string]world.Stat {
	return p.Entity.Stats()
}

func (p *AiPokemonEntity) Effects() []*world.Effect {
	return p.Entity.Effects()
}

func (p *AiPokemonEntity) SetEffects(effects []*world.Effect) {
	p.Entity.SetEffects(effects)
}

// AiPokemonEntity is intelligent

func (e *AiPokemonEntity) Script() world.AiScript {
	switch e.Profile.AggressionLevel {
	case AGGROpassive:
		return UpdateAggressiveAi
	case AGGROeager:
		return UpdatePassiveAi
	default:
		return nil
	}
}

func (p *AiPokemonEntity) Walking() *world.Position {
	return p.Entity.Walking()
}

func (p *AiPokemonEntity) SetWalking(walking *world.Position) {
	p.Entity.SetWalking(walking)
}

func (p *AiPokemonEntity) Abilities() map[string]*world.Ability {
	return p.Entity.Abilities()
}

func UpdateAggressiveAi(self world.Entity, now time.Time) {

}

func UpdatePassiveAi(self world.Entity, now time.Time) {

}
