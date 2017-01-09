package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/space"
	"strings"
	"time"
)

var vector_up = &space.Vector{0, -4}
var vector_right = &space.Vector{4, 0}
var vector_left = &space.Vector{-4, 0}
var vector_down = &space.Vector{0, 4}

type Movement struct {
	*space.Vector
	Lock bool
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*save.Player)
		m := &Movement{}
		p.AddPart(m)
	})
	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*save.Player)
		b := args[1].(*save.Binding)

		if strings.HasPrefix(b.SystemId, "walk") {
			p.Parts[part.MOVEMENT].(*Movement).Walk(b.SystemId[5:])
		}
	})
	event.On(event.BindingUp, func(args ...interface{}) {
		p := args[0].(*save.Player)
		b := args[1].(*save.Binding)

		if strings.HasPrefix(b.SystemId, "walk") {
			p.Parts[part.MOVEMENT].(*Movement).ClearWalk(b.SystemId[5:])
		}
	})
}

func (m *Movement) Part() string {
	return part.MOVEMENT
}

func (m *Movement) Update(u *save.Universe, e *save.Entity, d time.Duration) {
	if m.Lock {
		return
	}

	if m.Vector == nil {
		event.Fire(event.MovementUpdate, e, nil)
		return
	}

	e.Shape = e.Shape.Move(*m.Vector)
	event.Fire(event.MovementUpdate, e, m.Vector)

	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		if e.Id == entityId {
			continue
		}

		if space.DistanceShapeShape(e.Shape, entity.Shape) < 1 {
			event.Fire(event.MovementCollision, e, entity, u)
		}
	}
}

func (m *Movement) Snapshot() map[string]interface{} {
	return nil
}

func (m *Movement) Walk(direction string) {
	if m.Vector == nil {
		m.Vector = &space.Vector{}
	}

	if direction == "up" {
		v := m.Vector.Add(*vector_up)
		m.Vector = &v
	} else if direction == "right" {
		v := m.Vector.Add(*vector_right)
		m.Vector = &v
	} else if direction == "left" {
		v := m.Vector.Add(*vector_left)
		m.Vector = &v
	} else if direction == "down" {
		v := m.Vector.Add(*vector_down)
		m.Vector = &v
	}
}

func (m *Movement) ClearWalk(direction string) {
	if m.Vector == nil {
		return
	}

	if direction == "up" {
		v := m.Vector.Add(vector_up.Reverse())
		m.Vector = &v
	} else if direction == "right" {
		v := m.Vector.Add(vector_right.Reverse())
		m.Vector = &v
	} else if direction == "left" {
		v := m.Vector.Add(vector_left.Reverse())
		m.Vector = &v
	} else if direction == "down" {
		v := m.Vector.Add(vector_down.Reverse())
		m.Vector = &v
	}

	if m.Vector.Length() < 1 {
		m.Vector = nil
	}
}
