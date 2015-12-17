package entities

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type PokemonEntity struct {
	ospokemon.BasicPokemon
	PHYSICS     *world.Physics
	GRAPHICS    *world.Graphics
	ACTION      *world.Action
	ENTITYID    int
	CONTROL     uint8
	STATHANDLES map[string]world.Stat
	EFFECTS     []*world.Effect
	WALKING     *physics.Point
	ABILITIES   map[string]*world.Ability
}

// PokemonEntity is an entity

func (p *PokemonEntity) EntityId() int {
	return p.ENTITYID
}

func (p *PokemonEntity) SetEntityId(id int) {
	p.ENTITYID = id
}

func (p *PokemonEntity) Physics() *world.Physics {
	return p.PHYSICS
}

func (p *PokemonEntity) Graphics() *world.Graphics {
	return p.GRAPHICS
}

func (p *PokemonEntity) Action() *world.Action {
	return p.ACTION
}

func (p *PokemonEntity) SetAction(action *world.Action) {
	p.ACTION = action
}

// PokemonEntity is mortal

func (p *PokemonEntity) Control() uint8 {
	return p.CONTROL
}

func (p *PokemonEntity) SetControl(control uint8) {
	p.CONTROL = control
}

func (p *PokemonEntity) Stats() map[string]world.Stat {
	for key, _ := range p.BasicPokemon.STATS {
		if p.STATHANDLES[key] == nil {
			p.STATHANDLES[key] = &statHandle{key, p.BasicPokemon}
		}
	}

	return p.STATHANDLES
}

func (p *PokemonEntity) Effects() []*world.Effect {
	return p.EFFECTS
}

func (p *PokemonEntity) SetEffects(effects []*world.Effect) {
	p.EFFECTS = effects
}

// PokemonEntity is intelligent

func (p *PokemonEntity) Script() world.AiScript {
	return func(e world.Entity, now time.Time) {}
}

func (p *PokemonEntity) Walking() *physics.Point {
	return p.WALKING
}

func (p *PokemonEntity) SetWalking(walking *physics.Point) {
	p.WALKING = walking
}

func (p *PokemonEntity) Abilities() map[string]*world.Ability {
	return p.ABILITIES
}

// Trick (pronounced "hack") to store stats in one place: on the pokemon
type statHandle struct {
	Name    string
	Pokemon ospokemon.BasicPokemon
}

func (handle *statHandle) Value() int {
	return handle.Pokemon.Stats()[handle.Name].Value()
}

func (handle *statHandle) SetValue(value int) {
	handle.Pokemon.Stats()[handle.Name].SetValue(value)
}

func (handle *statHandle) MaxValue() int {
	return handle.Pokemon.Stats()[handle.Name].EffortValue()
}

func (handle *statHandle) SetMaxValue(maxvalue int) {
	handle.Pokemon.Stats()[handle.Name].SetEffortValue(maxvalue)
}

func (handle *statHandle) BaseMaxValue() int {
	return handle.Pokemon.Stats()[handle.Name].IndividualValue()
}

func (handle *statHandle) SetBaseMaxValue(basemaxvalue int) {
	handle.Pokemon.Stats()[handle.Name].SetIndividualValue(basemaxvalue)
}
