package engine

import (
	"github.com/ospokemon/ospokemon/physics"
	"math"
)

func CheckCollision(s1 physics.Shape, s2 physics.Shape) bool {
	return DistanceShapeShape(s1, s2) < 1
}

func DistanceShapeShape(s1 physics.Shape, s2 physics.Shape) float64 {
	switch shape1 := s1.(type) {
	case physics.Point:
		return distancePointShape(shape1, s2)
	case physics.Line:
		switch shape2 := s2.(type) {
		case physics.Point:
			return physics.DistancePointLine(shape2, shape1)
		case physics.Line:
			return physics.DistanceLineLine(shape1, shape2)
		case physics.Rect:
			return distanceLineRect(shape1, shape2)
		// case physics.Circle:
		//  return distanceLineCircle(shape1, shape2)
		default:
			return 42
		}
	case physics.Rect:
		switch shape2 := s2.(type) {
		case physics.Point:
			return distancePointRect(shape2, shape1)
		case physics.Line:
			return distanceLineRect(shape2, shape1)
		case physics.Rect:
			return distanceRectRect(shape1, shape2)
			// case physics.Circle:
			//  return distanceRectCircle(shape1, shape2) < 1
		default:
			return 42
		}
	case physics.Circle:
		switch shape2 := s2.(type) {
		case physics.Point:
			return distancePointCircle(shape2, shape1)
		// case physics.Line:
		//  return distanceLineCircle(shape2, shape1) < 1
		// case physics.Rect:
		//  return distanceRectCircle(shape2, shape1) < 1
		// case physics.Circle:
		//  return distanceCircleCircle(shape1, shape2) < 1
		default:
			return 42
		}
	default:
		return 42
	}
}

func distancePointShape(point physics.Point, shape physics.Shape) float64 {
	switch shape2 := shape.(type) {
	case physics.Point:
		return physics.DistancePointPoint(point, shape2)
	case physics.Line:
		return physics.DistancePointLine(point, shape2)
	case physics.Rect:
		return distancePointRect(point, shape2)
	case physics.Circle:
		return distancePointCircle(point, shape2)
	default:
		return 1
	}
}

func distancePointRect(point physics.Point, rect physics.Rect) float64 {
	var mindistance float64 = 1000

	for _, line := range rect.MakeLines() {
		distance := physics.DistancePointLine(point, line)

		if distance < mindistance {
			mindistance = distance
		}
	}

	// TODO check for point completely inside

	return mindistance
}

func distancePointCircle(point physics.Point, circle physics.Circle) float64 {
	distancetocenter := physics.DistancePointPoint(point, circle.Anchor)

	if distancetocenter <= circle.Radius {
		return 0
	} else {
		return distancetocenter - circle.Radius
	}
}

func distanceLineRect(line physics.Line, rect physics.Rect) float64 {
	var d float64 = 10000

	for _, line2 := range rect.MakeLines() {
		d = math.Min(d, physics.DistanceLineLine(line, line2))
	}

	return d
}

func distanceRectRect(rect1 physics.Rect, rect2 physics.Rect) float64 {
	var d float64 = 10000

	for _, line := range rect1.MakeLines() {
		d = math.Min(d, distanceLineRect(line, rect2))
	}

	return d
}
