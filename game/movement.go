package game

import (
	// "github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

var VECTup = &space.Vector{0, -1}
var VECTright = &space.Vector{1, 0}
var VECTleft = &space.Vector{-1, 0}
var VECTdown = &space.Vector{0, 1}

type Movement struct {
	Lock   bool
	Left   bool
	Up     bool
	Right  bool
	Down   bool
	Target *space.Point
}

type Walk string

func (m *Movement) Part() string {
	return part.Movement
}

func (w Walk) Part() string {
	return part.Walk
}

func (m *Movement) Update(u *Universe, e *Entity, d time.Duration) {
	if m.Lock {
		return
	}

	var vector space.Vector
	if m.Target != nil {
		vector = space.Line{
			e.Shape.Center(),
			*m.Target,
		}.Vector()
	} else {
		vector = m.makeWalkVector()
	}

	stats := e.Parts[part.Stats].(Stats)
	speedStat := stats["speed"]

	vector = vector.MakeUnit().Multiply(speedStat.Value)

	e.Move(vector, u)
}

func (m *Movement) Snapshot() map[string]interface{} {
	return nil
}

func (m *Movement) Walk(direction string) {
	if direction == "up" {
		m.Up = true
	} else if direction == "right" {
		m.Right = true
	} else if direction == "left" {
		m.Left = true
	} else if direction == "down" {
		m.Down = true
	}
}

func (m *Movement) ClearWalk(direction string) {
	if direction == "up" {
		m.Up = false
	} else if direction == "right" {
		m.Right = false
	} else if direction == "left" {
		m.Left = false
	} else if direction == "down" {
		m.Down = false
	}
}

func (m *Movement) makeWalkVector() space.Vector {
	vector := space.Vector{}
	if m.Left {
		vector = vector.Add(*VECTleft)
	}
	if m.Up {
		vector = vector.Add(*VECTup)
	}
	if m.Right {
		vector = vector.Add(*VECTright)
	}
	if m.Down {
		vector = vector.Add(*VECTdown)
	}
	return vector
}
