package objects

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/physics"
)

type Player struct {
	ENTITYID  int
	GRAPHIC   string
	COLLISION engine.Collision
	MAP       string
	SHAPE     physics.Shape
	ospokemon.BasicTrainer
	ACTION     *engine.Action
	CONTROL    engine.Control
	ABILITIES  []*engine.Ability
	STATS      map[string]*engine.Stat
	GRAPHICS   map[engine.AnimationType]string
	EFFECTS    []*engine.Effect
	WALKING    *physics.Point
	GAMEMASTER bool
	engine.Events
}

var Players = make(map[int]*Player)

var LoadPlayer func(id int) *Player

// Player is an Entity

func (p *Player) EntityId() *int {
	return &p.ENTITYID
}

func (p *Player) Graphic() *string {
	return &p.GRAPHIC
}

func (p *Player) Collision() *engine.Collision {
	return &p.COLLISION
}

func (p *Player) Map() *string {
	return &p.MAP
}

func (p *Player) Shape() physics.Shape {
	return p.SHAPE
}

func (p *Player) SetShape(shape physics.Shape) {
	p.SHAPE = shape
}

// Name taken care of

func (p *Player) Action() *engine.Action {
	return p.ACTION
}

func (p *Player) SetAction(action *engine.Action) {
	p.ACTION = action
}

func (p *Player) Control() *engine.Control {
	return &p.CONTROL
}

func (p *Player) Abilities() *[]*engine.Ability {
	return &p.ABILITIES
}

func (p *Player) Stats() map[string]*engine.Stat {
	return p.STATS
}

func (p *Player) Graphics() map[engine.AnimationType]string {
	return p.GRAPHICS
}

func (p *Player) Effects() *[]*engine.Effect {
	return &p.EFFECTS
}

func (p *Player) Walking() *physics.Point {
	return p.WALKING
}

func (p *Player) SetWalking(walking *physics.Point) {
	p.WALKING = walking
}

func (p *Player) IsGameMaster() bool {
	return p.GAMEMASTER
}

func GetPlayer(id int) *Player {
	if Players[id] == nil {
		Players[id] = LoadPlayer(id)
	}

	return Players[id]
}
