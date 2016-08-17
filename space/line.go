package space

type Line struct {
	P1 Point
	P2 Point
}

func (line Line) Copy() Line {
	return Line{
		P1: line.P1.Copy(),
		P2: line.P2.Copy(),
	}
}

func (line Line) Move(v Vector) Line {
	return Line{
		P1: line.P1.Move(v),
		P2: line.P2.Move(v),
	}
}

func (line Line) Vector() Vector {
	return Vector{
		DX: line.P2.X - line.P1.X,
		DY: line.P2.Y - line.P1.Y,
	}
}

func (line Line) Equation() func(float64) float64 {
	if line.P2.X-line.P1.X == 0 {
		return nil
	} else {
		return line.equation
	}
}

func (line Line) equation(x float64) float64 {
	m := (line.P2.Y - line.P1.Y) / (line.P2.X - line.P1.X)
	return (m * (x - line.P1.X)) + line.P1.Y
}
