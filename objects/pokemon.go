package objects

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/physics"
)

type PokemonEntity struct {
	ENTITYID  int
	GRAPHIC   string
	COLLISION engine.Collision
	SHAPE     physics.Shape
	ospokemon.BasicPokemon
	ACTION    *engine.Action
	CONTROL   engine.Control
	STATS     map[string]*engine.Stat
	ABILITIES []*engine.Ability
	GRAPHICS  map[engine.AnimationType]string
	EFFECTS   []*engine.Effect
	WALKING   *physics.Point
}

var Pokemon = make(map[int]*PokemonEntity)

// PokemonEntity is an entity

func (p *PokemonEntity) EntityId() *int {
	return &p.ENTITYID
}

func (p *PokemonEntity) Graphic() *string {
	return &p.GRAPHIC
}

func (p *PokemonEntity) Collision() *engine.Collision {
	return &p.COLLISION
}

func (p *PokemonEntity) Shape() physics.Shape {
	return p.SHAPE
}

func (p *PokemonEntity) SetShape(shape physics.Shape) {
	p.SHAPE = shape
}

func (p *PokemonEntity) Action() *engine.Action {
	return p.ACTION
}

func (p *PokemonEntity) SetAction(action *engine.Action) {
	p.ACTION = action
}

func (p *PokemonEntity) Control() *engine.Control {
	return &p.CONTROL
}

func (p *PokemonEntity) Abilities() *[]*engine.Ability {
	return &p.ABILITIES
}

func (p *PokemonEntity) Stats() map[string]*engine.Stat {
	return p.STATS
}

func (p *PokemonEntity) Graphics() map[engine.AnimationType]string {
	return p.GRAPHICS
}

func (p *PokemonEntity) Effects() *[]*engine.Effect {
	return &p.EFFECTS
}

func (p *PokemonEntity) Walking() *physics.Point {
	return p.WALKING
}

func (p *PokemonEntity) SetWalking(walking *physics.Point) {
	p.WALKING = walking
}
