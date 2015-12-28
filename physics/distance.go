package physics

import (
	"math"
)

func DistancePointPoint(point1 Point, point2 Point) float64 {
	dltx := point1.X - point2.X
	dlty := point1.Y - point2.Y
	return math.Sqrt((dlty * dlty) + (dltx * dltx))
}

func DistancePointLine(point Point, line Line) float64 {
	edgedist := math.Min(DistancePointPoint(point, line.P1), DistancePointPoint(point, line.P2))

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

	if line.Vector().DY == 0 {
		if point.X > math.Max(line.P1.X, line.P2.X) || point.X < math.Min(line.P1.X, line.P2.X) {
			return edgedist
		} else {
			return math.Abs(point.Y - line.P1.Y)
		}
	}

	x := pointsSlopesIntersect(line.Vector().AsSlope(), line.Vector().MakeRightAngle().AsSlope(), line.P1, point)

	if x >= math.Min(line.P1.X, line.P2.X) && x <= math.Max(line.P1.X, line.P2.X) {
		return DistancePointPoint(point, Point{x, equation(x)})
	} else {
		return edgedist
	}
}

func DistanceLineLine(line1 Line, line2 Line) float64 {
	edgedist := DistancePointLine(line1.P1, line2)
	edgedist = math.Min(edgedist, DistancePointLine(line1.P2, line2))
	edgedist = math.Min(edgedist, DistancePointLine(line2.P1, line1))
	edgedist = math.Min(edgedist, DistancePointLine(line2.P2, line1))

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

			return math.Min(DistancePointLine(Point{x, y}, line1), edgedist)
		}
	} else if eq2 == nil {
		x := line2.P1.X
		y := eq1(x)

		return math.Min(DistancePointLine(Point{x, y}, line2), edgedist)
	}

	if line1slope := line1.Vector().AsSlope(); line1slope == 0 {
		if line2slope := line2.Vector().AsSlope(); line2slope == 0 {
			return edgedist
		} else {
			y := line1.P1.Y
			x := (y - line2.P1.Y + line2slope*line2.P1.X) / line2slope

			return math.Min(DistancePointLine(Point{x, y}, line2), edgedist)
		}
	} else if line2.Vector().AsSlope() == 0 {
		y := line2.P1.Y
		x := (y - line1.P1.Y + line1slope*line1.P1.X) / line1slope

		return math.Min(DistancePointLine(Point{x, y}, line1), edgedist)
	}

	x := pointsSlopesIntersect(line1.Vector().AsSlope(), line2.Vector().AsSlope(), line1.P1, line2.P1)

	if x >= math.Min(line1.P1.X, line1.P2.X) && x >= math.Min(line2.P1.X, line2.P2.X) && x <= math.Max(line1.P1.X, line1.P2.X) && x <= math.Max(line2.P1.X, line2.P2.X) {
		return 0
	} else {
		return edgedist
	}
}

func pointsSlopesIntersect(m1 float64, m2 float64, p1 Point, p2 Point) float64 {
	y1 := p1.Y - m1*p1.X
	y2 := p2.Y - m2*p2.X
	return (y2 - y1) / (m1 - m2)
}
