package space

type PointShape Point

func (p PointShape) Center() Point {
	return Point(p)
}

func (p PointShape) Copy() Shape {
	return PointShape(Point(p).Copy())
}

func (p PointShape) Move(v Vector) Shape {
	return PointShape(Point(p).Move(v))
}
