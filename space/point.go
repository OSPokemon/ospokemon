package space

type Point struct {
	X float64
	Y float64
}

func (point Point) Copy() Point {
	return Point{
		X: point.X,
		Y: point.Y,
	}
}

func (p Point) Move(v Vector) Point {
	return Point{
		X: p.X + v.DX,
		Y: p.Y + v.DY,
	}
}

func (p Point) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"x": p.X,
		"y": p.Y,
	}
}
