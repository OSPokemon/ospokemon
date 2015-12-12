package world

import (
	"math"
)

func CreatePathVector(pos1 *Point, pos2 *Point, speed int) *Vector {
	ydiff := pos2.Y - pos1.Y
	xdiff := pos2.X - pos1.X
	mag := math.Sqrt(xdiff*xdiff + ydiff*ydiff)

	return &Vector{
		DX: xdiff / mag * float64(speed),
		DY: ydiff / mag * float64(speed),
	}
}
