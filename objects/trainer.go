package objects

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/physics"
)

type Trainer struct {
	ENTITYID  int
	GRAPHIC   string
	COLLISION engine.Collision
	MAP       string
	SHAPE     physics.Shape
	ospokemon.BasicTrainer
	ACTION    *engine.Action
	CONTROL   engine.Control
	ABILITIES []*engine.Ability
	STATS     map[string]*engine.Stat
	GRAPHICS  map[engine.AnimationType]string
	EFFECTS   []*engine.Effect
	WALKING   *physics.Point
	engine.Events
}

var Trainers = make(map[int]*Trainer)

var LoadTrainer func(id int) (*Trainer, error)
var CreateTrainer func(name string, class int) (*Trainer, error)
var SaveTrainer func(trainer *Trainer) error

// Trainer is an Entity

func (p *Trainer) EntityId() *int {
	return &p.ENTITYID
}

func (p *Trainer) Graphic() *string {
	return &p.GRAPHIC
}

func (p *Trainer) Collision() *engine.Collision {
	return &p.COLLISION
}

func (p *Trainer) Map() *string {
	return &p.MAP
}

func (p *Trainer) Shape() physics.Shape {
	return p.SHAPE
}

func (p *Trainer) SetShape(shape physics.Shape) {
	p.SHAPE = shape
}

// Name taken care of

func (p *Trainer) Action() *engine.Action {
	return p.ACTION
}

func (p *Trainer) SetAction(action *engine.Action) {
	p.ACTION = action
}

func (p *Trainer) Control() *engine.Control {
	return &p.CONTROL
}

func (p *Trainer) Abilities() *[]*engine.Ability {
	return &p.ABILITIES
}

func (p *Trainer) Stats() map[string]*engine.Stat {
	return p.STATS
}

func (p *Trainer) Graphics() map[engine.AnimationType]string {
	return p.GRAPHICS
}

func (p *Trainer) Effects() *[]*engine.Effect {
	return &p.EFFECTS
}

func (p *Trainer) Walking() *physics.Point {
	return p.WALKING
}

func (p *Trainer) SetWalking(walking *physics.Point) {
	p.WALKING = walking
}

func GetTrainer(id int) *Trainer {
	if Trainers[id] == nil {
		if trainer, err := LoadTrainer(id); err == nil {
			Trainers[id] = trainer
		} else {
			log.WithFields(log.Fields{
				"TrainerId": id,
				"Error":     err.Error(),
			}).Info("Trainer lookup failed")
		}
	}

	return Trainers[id]
}
