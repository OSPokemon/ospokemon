package engine

import (
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
)

const EVNT_SpaceDivide = "engine/Space.Divide"

type Space struct {
	space.Rect
	Division *space.Line
	Sub      *[2]*Space
	Entities []Entity
	util.Eventer
}

func MakeSpace() *Space {
	return &Space{
		Rect: space.Rect{
			Anchor:    space.Point{},
			Dimension: space.Vector{},
		},
		Division: nil,
		Sub:      nil,
		Entities: make([]Entity, 0),
		Eventer:  make(util.Eventer),
	}
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

	s.Sub[0] = MakeSpace()
	s.Sub[0].Rect.Anchor = s.Anchor.Copy()
	s.Sub[0].Rect.Dimension.DX = s.Division.P2.X - s.Anchor.X
	s.Sub[0].Rect.Dimension.DY = s.Division.P2.Y - s.Anchor.Y

	s.Sub[1] = MakeSpace()
	s.Sub[1].Rect.Anchor = s.Division.P1.Copy()
	s.Sub[1].Rect.Dimension.DX = s.Division.P2.X - s.Anchor.X
	s.Sub[1].Rect.Dimension.DY = s.Division.P2.Y - s.Anchor.Y
}

func (s *Space) assignDivision() {
	// TODO
}
