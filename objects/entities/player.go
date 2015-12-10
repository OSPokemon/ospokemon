package entities

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type Player struct {
	ospokemon.BasicTrainer
	PHYSICS   *world.Physics
	GRAPHICS  *world.Graphics
	ACTION    *world.Action
	ENTITYID  int
	CONTROL   uint8
	STATS     map[string]world.Stat
	EFFECTS   []*world.Effect
	WALKING   *world.Position
	ABILITIES map[string]*world.Ability
	world.Events
}

// Player is an Entity

func (p *Player) EntityId() int {
	return p.ENTITYID
}

func (p *Player) SetEntityId(id int) {
	p.ENTITYID = id
}

func (p *Player) Physics() *world.Physics {
	return p.PHYSICS
}

func (p *Player) Graphics() *world.Graphics {
	return p.GRAPHICS
}

func (p *Player) Action() *world.Action {
	return p.ACTION
}

func (p *Player) SetAction(action *world.Action) {
	p.ACTION = action
}

// Player is mortal

func (p *Player) Control() uint8 {
	return p.CONTROL
}

func (p *Player) SetControl(control uint8) {
	p.CONTROL = control
}

func (p *Player) Stats() map[string]world.Stat {
	return p.STATS
}

func (p *Player) Effects() []*world.Effect {
	return p.EFFECTS
}

func (p *Player) SetEffects(effects []*world.Effect) {
	p.EFFECTS = effects
}

// Player is intelligent

func (p *Player) Script() world.AiScript {
	return func(e world.Entity, now time.Time) {}
}

func (p *Player) Walking() *world.Position {
	return p.WALKING
}

func (p *Player) SetWalking(walking *world.Position) {
	p.WALKING = walking
}

func (p *Player) Abilities() map[string]*world.Ability {
	return p.ABILITIES
}

type PlayerStat struct {
	VALUE        int
	MAXVALUE     int
	BASEMAXVALUE int
}

func (stat *PlayerStat) Value() int {
	return stat.VALUE
}

func (stat *PlayerStat) SetValue(value int) {
	stat.VALUE = value
}

func (stat *PlayerStat) MaxValue() int {
	return stat.MAXVALUE
}

func (stat *PlayerStat) SetMaxValue(maxvalue int) {
	stat.MAXVALUE = maxvalue
}

func (stat *PlayerStat) BaseMaxValue() int {
	return stat.BASEMAXVALUE
}

func (stat *PlayerStat) SetBaseMaxValue(basemaxvalue int) {
	stat.BASEMAXVALUE = basemaxvalue
}
