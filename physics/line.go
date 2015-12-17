package physics

type Line struct {
	P1 Point
	P2 Point
}

func (line Line) Copy() Shape {
	return Line{
		P1: line.P1.Copy().(Point),
		P2: line.P2.Copy().(Point),
	}
}

func (line Line) Move(v Vector) Shape {
	return Line{
		P1: line.P1.Move(v).(Point),
		P2: line.P2.Move(v).(Point),
	}
}

func (line Line) Vector() Vector {
	return Vector{
		DX: line.P2.X - line.P1.X,
		DY: line.P2.Y - line.P1.Y,
	}
}

func YIntersect(slope float64, point Point) float64 {
	return point.Y - slope*point.X
}
