package physics

type Rect struct {
	Anchor   Point
	Rotation Vector
	Height   float64
	Width    float64
}

func (r Rect) Copy() Shape {
	return Rect{
		Anchor:   r.Anchor.Copy().(Point),
		Rotation: r.Rotation.Copy(),
		Height:   r.Height,
		Width:    r.Width,
	}
}

func (r Rect) Move(v Vector) Shape {
	return Rect{
		Anchor:   r.Anchor.Move(v).(Point),
		Rotation: r.Rotation.Copy(),
		Height:   r.Height,
		Width:    r.Width,
	}
}

func (r Rect) MakeLines() []Line {
	vectorstraight := r.Rotation.MakeUnit()
	vectoracross := r.Rotation.MakeRightAngle().MakeUnit()

	point1 := r.Anchor
	point2 := r.Anchor.Move(vectorstraight.Multiply(r.Height)).(Point)
	point3 := r.Anchor.Move(vectoracross.Multiply(r.Width)).(Point)
	point4 := point2.Move(vectoracross.Multiply(r.Width)).(Point)

	return []Line{
		Line{point1, point2},
		Line{point1, point3},
		Line{point2, point4},
		Line{point3, point4},
	}
}
