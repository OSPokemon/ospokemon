package space

import (
	"math"
)

func distancePointPoint(point1 Point, point2 Point) float64 {
	dltx := point1.X - point2.X
	dlty := point1.Y - point2.Y
	return math.Sqrt((dlty * dlty) + (dltx * dltx))
}

func distancePointLine(point Point, line Line) float64 {
	edgedist := math.Min(distancePointPoint(point, line.P1), distancePointPoint(point, line.P2))

	if edgedist == 0 {
		return 0
	}

	equation := line.Equation()

	if equation == nil {
		if point.Y > math.Max(line.P1.Y, line.P2.Y) || point.Y < math.Min(line.P1.Y, line.P2.Y) {
			return edgedist
		} else {
			return math.Abs(point.X - line.P1.X)
		}
	}
	if equation(point.X) == point.Y {
		if point.X > math.Max(line.P1.X, line.P2.X) || point.X < math.Min(line.P1.X, line.P2.X) {
			return edgedist
		} else {
			return 0
		}
	}

	vector := line.Vector()

	if vector.DY == 0 {
		if point.X > math.Max(line.P1.X, line.P2.X) || point.X < math.Min(line.P1.X, line.P2.X) {
			return edgedist
		} else {
			return math.Abs(point.Y - line.P1.Y)
		}
	}

	x := pointsSlopesIntersect(vector.AsSlope(), vector.MakeRightAngle().AsSlope(), line.P1, point)

	if x >= math.Min(line.P1.X, line.P2.X) && x <= math.Max(line.P1.X, line.P2.X) {
		return distancePointPoint(point, Point{x, equation(x)})
	} else {
		return edgedist
	}
}

func distancePointRect(point Point, rect Rect) float64 {
	corner1 := rect.Anchor
	corner2 := rect.Anchor.Move(rect.Dimension)

	if math.Min(corner1.X, corner2.X) < point.X && math.Max(corner1.X, corner2.X) > point.X && math.Min(corner1.Y, corner2.Y) < point.Y && math.Max(corner1.Y, corner2.Y) > point.Y {
		return 0
	}

	var mindistance float64 = 1000

	for _, line := range rect.MakeLines() {
		distance := distancePointLine(point, line)

		if distance < mindistance {
			mindistance = distance
		}
	}

	return mindistance
}

func distancePointCircle(point Point, circle Circle) float64 {
	return distancePointPoint(point, circle.Anchor) - circle.Radius
}

func distanceLineLine(line1 Line, line2 Line) float64 {
	edgedist := distancePointLine(line1.P1, line2)
	edgedist = math.Min(edgedist, distancePointLine(line1.P2, line2))
	edgedist = math.Min(edgedist, distancePointLine(line2.P1, line1))
	edgedist = math.Min(edgedist, distancePointLine(line2.P2, line1))

	if edgedist == 0 {
		return 0
	}

	eq1 := line1.Equation()
	eq2 := line2.Equation()

	if eq1 == nil {
		if eq2 == nil {
			return edgedist
		} else {
			x := line1.P1.X
			y := eq2(x)

			if y >= math.Min(line1.P1.Y, line1.P2.Y) && y <= math.Max(line1.P1.Y, line1.P2.Y) {
				edgedist = math.Min(edgedist, distancePointLine(Point{x, y}, line2))
			}

			return edgedist
		}
	} else if eq2 == nil {
		x := line2.P1.X
		y := eq1(x)

		if y >= math.Min(line2.P1.Y, line2.P2.Y) && y <= math.Max(line2.P1.Y, line2.P2.Y) {
			edgedist = math.Min(edgedist, distancePointLine(Point{x, y}, line1))
		}

		return edgedist
	}

	if line1slope := line1.Vector().AsSlope(); line1slope == 0 {
		if line2slope := line2.Vector().AsSlope(); line2slope == 0 {
			return edgedist
		} else {
			y := line1.P1.Y
			x := (y - line2.P1.Y + line2slope*line2.P1.X) / line2slope

			if x >= math.Min(line1.P1.X, line1.P2.X) && x <= math.Max(line1.P1.X, line1.P2.X) {
				edgedist = math.Min(edgedist, distancePointLine(Point{x, y}, line2))
			}

			return edgedist
		}
	} else if line2.Vector().AsSlope() == 0 {
		y := line2.P1.Y
		x := (y - line1.P1.Y + line1slope*line1.P1.X) / line1slope

		if x >= math.Min(line2.P1.X, line2.P2.X) && x <= math.Max(line2.P1.X, line2.P2.X) {
			edgedist = math.Min(edgedist, distancePointLine(Point{x, y}, line1))
		}

		return edgedist
	}

	x := pointsSlopesIntersect(line1.Vector().AsSlope(), line2.Vector().AsSlope(), line1.P1, line2.P1)

	if x >= math.Min(line1.P1.X, line1.P2.X) && x >= math.Min(line2.P1.X, line2.P2.X) && x <= math.Max(line1.P1.X, line1.P2.X) && x <= math.Max(line2.P1.X, line2.P2.X) {
		return 0
	} else {
		return edgedist
	}
}

func distanceLineRect(line Line, rect Rect) float64 {
	var d float64 = 10000

	for _, line2 := range rect.MakeLines() {
		d = math.Min(d, distanceLineLine(line, line2))
	}

	return d
}

func distanceLineCircle(line Line, circle Circle) float64 {
	return distancePointLine(circle.Anchor, line) - circle.Radius
}

func distanceRectRect(rect1 Rect, rect2 Rect) float64 {
	var d float64 = 10000

	for _, point := range rect1.MakeCorners() {
		d = math.Min(d, distancePointRect(point, rect2))

		if d < 1 {
			return d
		}
	}
	for _, point := range rect2.MakeCorners() {
		d = math.Min(d, distancePointRect(point, rect1))

		if d < 1 {
			return d
		}
	}

	for _, line := range rect1.MakeLines() {
		d = math.Min(d, distanceLineRect(line, rect2))

		if d < 1 {
			return d
		}
	}

	return d
}

func distanceRectCircle(rect Rect, circle Circle) float64 {
	return distancePointRect(circle.Anchor, rect) - circle.Radius
}

func distanceCircleCircle(circle1 Circle, circle2 Circle) float64 {
	return distancePointPoint(circle1.Anchor, circle2.Anchor) - circle1.Radius - circle2.Radius
}

func distancePointShape(point Point, shape Shape) float64 {
	switch shape2 := shape.(type) {
	case *Rect:
		return distancePointRect(point, *shape2)
	case *Circle:
		return distancePointCircle(point, *shape2)
	default:
		return 1
	}
}

func DistanceShapeShape(s1 Shape, s2 Shape) float64 {
	switch shape1 := s1.(type) {
	case *Rect:
		switch shape2 := s2.(type) {
		case *Rect:
			return distanceRectRect(*shape1, *shape2)
		case *Circle:
			return distanceRectCircle(*shape1, *shape2)
		default:
			return 42
		}
	case *Circle:
		switch shape2 := s2.(type) {
		case *Rect:
			return distanceRectCircle(*shape2, *shape1)
		case *Circle:
			return distanceCircleCircle(*shape1, *shape2)
		default:
			return 42
		}
	default:
		return 42
	}
}

func pointsSlopesIntersect(m1 float64, m2 float64, p1 Point, p2 Point) float64 {
	y1 := p1.Y - m1*p1.X
	y2 := p2.Y - m2*p2.X
	return (y2 - y1) / (m1 - m2)
}
