package physics

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (point Point) Copy() Shape {
	return Point{
		X: point.X,
		Y: point.Y,
	}
}

func (p Point) Move(v Vector) Shape {
	return Point{
		X: p.X + v.DX,
		Y: p.Y + v.DY,
	}
}

func DistancePointPoint(point1 Point, point2 Point) float64 {
	dltx := point1.X - point2.X
	dlty := point1.Y - point2.Y
	return math.Sqrt((dlty * dlty) + (dltx * dltx))
}
