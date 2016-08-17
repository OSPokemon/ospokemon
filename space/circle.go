package space

type Circle struct {
	Anchor Point
	Radius float64
}

func (c Circle) Center() Point {
	return c.Anchor.Copy()
}

func (c Circle) Copy() Shape {
	return Circle{
		Anchor: c.Anchor.Copy(),
		Radius: c.Radius,
	}
}

func (c Circle) Move(v Vector) Shape {
	return Circle{
		Anchor: c.Anchor.Move(v),
		Radius: c.Radius,
	}
}
