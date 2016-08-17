package space

type Rect struct {
	Anchor    Point
	Dimension Vector
}

func (r Rect) Center() Point {
	vectorx := Vector{DX: r.Dimension.DX / 2}
	vectory := Vector{DY: r.Dimension.DY / 2}
	return r.Anchor.Move(vectorx).Move(vectory)
}

func (r Rect) Copy() Shape {
	return Rect{
		Anchor:    r.Anchor.Copy(),
		Dimension: r.Dimension.Copy(),
	}
}

func (r Rect) Move(v Vector) Shape {
	return Rect{
		Anchor:    r.Anchor.Move(v),
		Dimension: r.Dimension,
	}
}

func (r Rect) MakeCorners() []Point {
	vectorx := Vector{DX: r.Dimension.DX}
	vectory := Vector{DY: r.Dimension.DY}

	return []Point{
		r.Anchor,
		r.Anchor.Move(vectorx),
		r.Anchor.Move(vectory),
		r.Anchor.Move(vectorx).Move(vectory),
	}
}

func (r Rect) MakeLines() []Line {
	points := r.MakeCorners()

	return []Line{
		Line{points[0], points[1]},
		Line{points[0], points[2]},
		Line{points[1], points[3]},
		Line{points[2], points[3]},
	}
}
