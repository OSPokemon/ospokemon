package physics

import (
	"testing"
)

func TestPointCopy(test *testing.T) {
	point1 := Point{350, 200}
	point2 := point1.Copy().(Point)

	if point1.X != point2.X {
		test.Fail()
	}
	if point1.Y != point2.Y {
		test.Fail()
	}
}

func TestPointMove(test *testing.T) {
	point1 := Point{350, 200}
	point2 := point1.Move(Vector{100, 200}).(Point)

	if point1.X+100 != point2.X {
		test.Fail()
	}
	if point1.Y+200 != point2.Y {
		test.Fail()
	}
}
