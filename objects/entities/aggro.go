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

func (p *AiPokemonEntity) Name() string {
	return p.Entity.BasicPokemon.Name()
}

func (p *AiPokemonEntity) Physics() *world.Physics {
	return p.Entity.PHYSICS
}

func (p *AiPokemonEntity) Graphics() *world.Graphics {
	return p.Entity.GRAPHICS
}

func (p *AiPokemonEntity) Controls() *world.Controls {
	return p.Entity.CONTROLS
}

func (p *AiPokemonEntity) Effects() world.Effects {
	return p.Entity.EFFECTS
}

func (p *AiPokemonEntity) SetEffects(effects world.Effects) {
	p.Entity.EFFECTS = effects
}

func (p *AiPokemonEntity) Health() int {
	return p.Entity.BasicPokemon.Stats()["health"].Value()
}

func (p *AiPokemonEntity) MaxHealth() int {
	// TODO load species health
	return 1000
}

func (p *AiPokemonEntity) SetHealth(health int) {
	p.Entity.BasicPokemon.Stats()["health"].SetValue(health)
}

func (p *AiPokemonEntity) Speed() int {
	return p.Entity.BasicPokemon.Stats()["speed"].Value()
}

func UpdateAggressiveAi(self world.Entity, now time.Time) {

}

func UpdatePassiveAi(self world.Entity, now time.Time) {

}
