package update

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/world"
	"math"
)

func MoveEntity(entity world.Entity, v physics.Vector) {
	if mortal, ok := entity.(world.Mortality); ok && world.IsStuck(mortal) {
		return
	}

	nextShape := entity.Physics().Shape.Move(v)

	if entity.Physics().Solid {
		for _, entity2 := range world.Entities {
			if entity == entity2 {
				continue
			}
			if !entity2.Physics().Solid {
				continue
			}

			if CheckCollision(nextShape, entity2.Physics().Shape) {
				return
			}
		}
	}

	log.WithFields(log.Fields{
		"Entity": entity.Name(),
		"Vector": v,
		"Shape":  nextShape,
	}).Debug("Move entity")

	entity.Physics().Shape = nextShape
}

func CheckCollision(s1 physics.Shape, s2 physics.Shape) bool {
	switch shape1 := s1.(type) {
	case physics.Point:
		return distancePointShape(shape1, s2) < 1
	case physics.Line:
		switch shape2 := s2.(type) {
		case physics.Point:
			return distancePointLine(shape2, shape1) < 1
		// case physics.Line:
		// 	return distanceLineLine(shape1, shape2)
		case physics.Rect:
			return distanceLineRect(shape1, shape2) < 1
		// case physics.Circle:
		// 	return distanceLineCircle(shape1, shape2)
		default:
			return false
		}
	case physics.Rect:
		switch shape2 := s2.(type) {
		case physics.Point:
			return distancePointRect(shape2, shape1) < 1
		case physics.Line:
			return distanceLineRect(shape2, shape1) < 1
		case physics.Rect:
			return distanceRectRect(shape1, shape2) < 1
			// case physics.Circle:
			// 	return distanceRectCircle(shape1, shape2) < 1
		default:
			return false
		}
	case physics.Circle:
		switch shape2 := s2.(type) {
		case physics.Point:
			return distancePointCircle(shape2, shape1) < 1
		// case physics.Line:
		// 	return distanceLineCircle(shape2, shape1) < 1
		// case physics.Rect:
		// 	return distanceRectCircle(shape2, shape1) < 1
		// case physics.Circle:
		// 	return distanceCircleCircle(shape1, shape2) < 1
		default:
			return false
		}
	default:
		return false
	}
}

func distancePointShape(point physics.Point, shape physics.Shape) float64 {
	switch shape2 := shape.(type) {
	case physics.Point:
		return physics.DistancePointPoint(point, shape2)
	case physics.Line:
		return distancePointLine(point, shape2)
	case physics.Rect:
		return distancePointRect(point, shape2)
	case physics.Circle:
		return distancePointCircle(point, shape2)
	default:
		return 1
	}
}

func distancePointLine(point physics.Point, line physics.Line) float64 {
	if math.Floor(line.P1.X) == math.Floor(point.X) && math.Floor(line.P1.Y) == math.Floor(point.Y) {
		return 0
	}
	if math.Floor(line.P2.X) == math.Floor(point.X) && math.Floor(line.P2.Y) == math.Floor(point.Y) {
		return 0
	}

	m1 := line.Vector().AsSlope()         // slope of the main line
	y1 := physics.YIntersect(m1, line.P1) // y intersect of the main line

	m2 := line.Vector().MakeRightAngle().AsSlope() // slope of the right-angle intersection line
	y2 := physics.YIntersect(m2, point)            // y intersect using p and m2

	xint := (y2 - y1) / (m1 - m2)

	if (xint <= line.P1.X && xint >= line.P2.X) || (xint >= line.P1.X && xint <= line.P2.X) {
		yint := m1*xint + y1 // mx + b

		intersectpoint := physics.Point{
			X: xint,
			Y: yint,
		}

		return physics.DistancePointPoint(intersectpoint, point)
	} else {
		return math.Min(physics.DistancePointPoint(point, line.P1), physics.DistancePointPoint(point, line.P2))
	}
}

func distancePointRect(point physics.Point, rect physics.Rect) float64 {
	var mindistance float64 = 1000

	for _, line := range rect.MakeLines() {
		distance := distancePointLine(point, line)

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

func distanceLineLine(line1 physics.Line, line2 physics.Line) float64 {
	d := math.Min(distancePointLine(line1.P1, line2), distancePointLine(line1.P2, line2))
	d = math.Min(d, distancePointLine(line2.P1, line1))
	d = math.Min(d, distancePointLine(line2.P2, line1))

	return d
}

func distanceLineRect(line physics.Line, rect physics.Rect) float64 {
	var d float64 = 10000

	for _, line2 := range rect.MakeLines() {
		d = math.Min(d, distanceLineLine(line, line2))
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
