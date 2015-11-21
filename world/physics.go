package world

import (
	"math"
)

type Position struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
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
	return math.Sqrt((dlty * dlty) + (dltx * dltx))
}

func (p *Position) Add(v *Vector) Position {
	return Position{
		X: p.X + (math.Sin(v.Radians) * v.Distance),
		Y: p.Y + (math.Cos(v.Radians) * v.Distance),
	}
}

func (phys1 *Physics) CheckCollision(phys2 *Physics) bool {
	p1l := phys1.X
	p1r := p1l + phys1.Width
	p1t := phys1.Y
	p1b := p1t + phys1.Height

	p2l := phys2.X
	p2r := p2l + phys2.Width
	p2t := phys2.Y
	p2b := p2t + phys2.Height

	return p1l < p2r && p1r > p2l && p1t < p2b && p1b > p2t
}
