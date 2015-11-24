package world

import (
	"math"
)

type Position struct {
	X float64
	Y float64
}

type Size struct {
	Width  int
	Height int
}

type Physics struct {
	Position
	Size
	Solid   bool
	Walking *Position
}

type Vector struct {
	DX float64
	DY float64
}

func GetDistance(p1 *Position, p2 *Position) float64 {
	dltx := p1.X - p2.X
	dlty := p1.Y - p2.Y
	return math.Sqrt(float64(dlty*dlty) + float64(dltx*dltx))
}

func (p *Position) Add(v *Vector) Position {
	return Position{
		X: p.X + v.DX,
		Y: p.Y + v.DY,
	}
}

func (phys1 *Physics) CheckCollision(phys2 *Physics) bool {
	p1l := phys1.X
	p1r := p1l + float64(phys1.Width)
	p1t := phys1.Y
	p1b := p1t + float64(phys1.Height)

	p2l := phys2.X
	p2r := p2l + float64(phys2.Width)
	p2t := phys2.Y
	p2b := p2t + float64(phys2.Height)

	return p1l < p2r && p1r > p2l && p1t < p2b && p1b > p2t
}
