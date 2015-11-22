package data

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/world"
)

type Player struct {
	ospokemon.BasicTrainer
	HEALTH    int
	MAXHEALTH int
	SPEED     int
	PHYSICS   *world.Physics
	GRAPHICS  *world.Graphics
	CONTROLS  *world.Controls
	EFFECTS   world.Effects
}

func (p *Player) Physics() *world.Physics {
	return p.PHYSICS
}

func (p *Player) Graphics() *world.Graphics {
	return p.GRAPHICS
}

func (p *Player) Controls() *world.Controls {
	return p.CONTROLS
}

func (p *Player) Effects() world.Effects {
	return p.EFFECTS
}

func (p *Player) SetEffects(effects world.Effects) {
	p.EFFECTS = effects
}

func (p *Player) Health() int {
	return p.HEALTH
}

func (p *Player) MaxHealth() int {
	return p.MAXHEALTH
}

func (p *Player) SetHealth(health int) {
	p.HEALTH = health
}

func (p *Player) Speed() int {
	return p.SPEED
}
