package data

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/world"
)

type PokemonEntity struct {
	ospokemon.BasicPokemon
	PHYSICS  *world.Physics
	GRAPHICS *world.Graphics
	CONTROLS *world.Controls
	EFFECTS  world.Effects
}

func (p *PokemonEntity) Physics() *world.Physics {
	return p.PHYSICS
}

func (p *PokemonEntity) Graphics() *world.Graphics {
	return p.GRAPHICS
}

func (p *PokemonEntity) Controls() *world.Controls {
	return p.CONTROLS
}

func (p *PokemonEntity) Effects() world.Effects {
	return p.EFFECTS
}

func (p *PokemonEntity) SetEffects(effects world.Effects) {
	p.EFFECTS = effects
}

func (p *PokemonEntity) Health() int {
	return p.BasicPokemon.Stats()["health"].Value()
}

func (p *PokemonEntity) MaxHealth() int {
	// TODO load species health
	return 1000
}

func (p *PokemonEntity) SetHealth(health int) {
	p.BasicPokemon.Stats()["health"].SetValue(health)
}

func (p *PokemonEntity) Speed() int {
	return p.BasicPokemon.Stats()["speed"].Value()
}
