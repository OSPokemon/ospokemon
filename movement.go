package ospokemon

import (
	"ospokemon.com/space"
	"time"
)

const PARTmovement = "movement"
const PARTwalk = "walk"

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
	return PARTmovement
}

func (parts Parts) GetMovement() *Movement {
	movement, _ := parts[PARTmovement].(*Movement)
	return movement
}

func (w Walk) Part() string {
	return PARTwalk
}

func (parts Parts) GetWalk() Walk {
	walk, _ := parts[PARTwalk].(Walk)
	return walk
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

	stats := e.GetStats()
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
	} else {
		return
	}

	m.Target = nil
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
	} else {
		return
	}

	m.Target = nil
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
