package physics

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
