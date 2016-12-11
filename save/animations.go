package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/space"
	"math"
	"time"
)

const COMP_Animations = "animations"

type Animations map[string]string

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		animations := make(Animations)
		p.Entity.AddComponent(animations)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		animations := p.Entity.Component(COMP_Animations).(Animations)

		if c, err := GetClass(p.Class); c != nil {
			animations.clear()
			for key, value := range c.Animations {
				animations[key] = value
			}
			p.Entity.Image = c.Animations["portrait"]
		} else {
			logrus.WithFields(logrus.Fields{
				"Username": p.Username,
				"Class":    p.Class,
			}).Warn("save.Animations: " + err.Error())
		}
	})

	event.On(event.MovementUpdate, func(args ...interface{}) {
		e := args[0].(*Entity)
		v, ok := args[1].(*space.Vector)

		animations, ok := e.Component(COMP_Animations).(Animations)

		if !ok {
			return
		}

		if v == nil {
			e.Image = animations["portrait"]
		} else if slope := v.AsSlope(); slope == math.Inf(-1) {
			e.Image = animations["walk-up"]
		} else if slope == math.Inf(1) {
			e.Image = animations["walk-down"]
		} else if v.DX > 0 {
			e.Image = animations["walk-right"]
		} else {
			e.Image = animations["walk-left"]
		}
	})
}

func (a Animations) clear() {
	for k := range a {
		delete(a, k)
	}
}

func (a Animations) copy() {
}

func (a Animations) Id() string {
	return COMP_Animations
}

func (a Animations) Update(u *Universe, e *Entity, d time.Duration) {
	// nothing
}

func (a Animations) Snapshot() map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range a {
		m[k] = v
	}

	return m
}
