package world

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
)

func CreatePathVector(shape physics.Shape, point physics.Point, speed int) physics.Vector {
	pos := getAnchorPoint(shape)

	ydiff := point.Y - pos.Y
	xdiff := point.X - pos.X

	vector := physics.Vector{
		DY: ydiff,
		DX: xdiff,
	}

	log.WithFields(log.Fields{
		"Anchor":      pos,
		"Destination": point,
		"Vector":      vector,
		"Speed":       speed,
	}).Debug("Create path vector")

	vector = vector.MakeUnit().Multiply(float64(speed))

	return vector
}

func getAnchorPoint(s physics.Shape) physics.Point {
	var point physics.Point

	switch shape := s.(type) {
	case physics.Point:
		point = shape
		break
	case physics.Line:
		point = shape.P1
		break
	case physics.Circle:
		point = shape.Anchor
		break
	case physics.Rect:
		// halfrotation := shape.Rotation.Multiply(shape.Width / 2)
		// halfdownward := halfrotation.MakeRightAngle()

		// point = shape.Anchor.Move(halfrotation).Move(halfdownward).(physics.Point)
		point = shape.Anchor
		break
	}

	return point
}
