package world

import (
	"math"
)

type Position struct {
	X int64
	Y int64
}

type Size struct {
	Width  int
	Height int
}

type Physics struct {
	Position
	Size
}

type Vector struct {
	Radians  float64
	Distance float64
}

func GetDistance(p1 *Position, p2 *Position) float64 {
	dltx := p2.X - p1.X
	dlty := p2.Y - p1.Y
	return math.Sqrt(float64(dlty*dlty) + float64(dltx*dltx))
}

func (p *Position) Add(v *Vector) Position {
	return Position{
		X: p.X + int64(math.Sin(v.Radians)*v.Distance),
		Y: p.Y + int64(math.Cos(v.Radians)*v.Distance),
	}
}

func (phys1 *Physics) CheckCollision(phys2 *Physics) bool {
	p1l := phys1.X
	p1r := p1l + int64(phys1.Width)
	p1t := phys1.Y
	p1b := p1t + int64(phys1.Height)

	p2l := phys2.X
	p2r := p2l + int64(phys2.Width)
	p2t := phys2.Y
	p2b := p2t + int64(phys2.Height)

	return p1l < p2r && p1r > p2l && p1t < p2b && p1b > p2t
}
