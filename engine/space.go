package engine

import (
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
)

const EVNT_SpaceDivide = "ospokemon/engine/Space.Divide"

type Space struct {
	Name string
	space.Rect
	Division *space.Line
	Sub      *[2]*Space
	Entities []Entity
	util.Eventer
}

func (s *Space) Divide() {
	if s.Division != nil {
		s.Sub[0].Divide()
		s.Sub[1].Divide()
		return
	}

	s.createDivision()
	s.assignDivision()

	util.Event.Fire(EVNT_SpaceDivide, s)
	s.Fire(EVNT_SpaceDivide, s)
}

func (s *Space) Capacity() uint {
	cap := uint(len(s.Entities))

	if s.Division != nil {
		cap += s.Sub[0].Capacity() + s.Sub[1].Capacity()
	}

	return cap
}

func (s *Space) createDivision() {
	s.Division = &space.Line{s.Anchor.Copy(), s.Anchor.Move(s.Dimension)}

	if s.Dimension.DX > s.Dimension.DY {
		halfx := s.Dimension.DX / 2
		s.Division.P1.X += halfx
		s.Division.P2.X -= halfx
	} else {
		halfy := s.Dimension.DY / 2
		s.Division.P1.Y += halfy
		s.Division.P2.Y -= halfy
	}

	s.Sub = &[2]*Space{}

	s.Sub[0] = &Space{
		Name: s.Name + ".1",
		Rect: space.Rect{
			Anchor: s.Anchor.Copy(),
			Dimension: space.Vector{
				DX: s.Division.P2.X - s.Anchor.X,
				DY: s.Division.P2.Y - s.Anchor.Y,
			},
		},
		Division: nil,
		Sub:      nil,
		Entities: make([]Entity, 0),
	}

	s.Sub[1] = &Space{
		Name: s.Name + ".2",
		Rect: space.Rect{
			Anchor: s.Division.P1.Copy(),
			Dimension: space.Vector{
				DX: s.Division.P2.X - s.Anchor.X,
				DY: s.Division.P2.Y - s.Anchor.Y,
			},
		},
		Division: nil,
		Sub:      nil,
		Entities: make([]Entity, 0),
	}
}

func (s *Space) assignDivision() {
	// TODO
}
