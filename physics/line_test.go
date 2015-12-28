package physics

import (
	// "math"
	"testing"
)

func TestLineCopy(test *testing.T) {
	line1 := Line{Point{0, 0}, Point{100, 100}}
	line2 := line1.Copy().(Line)

	if line2.P1.X != 0 {
		test.Error("line copy point 1 x")
	}
	if line2.P1.Y != 0 {
		test.Error("line copy point 1 y")
	}
	if line2.P2.X != 100 {
		test.Error("line copy point 2 x")
	}
	if line2.P2.Y != 100 {
		test.Error("line copy point 2 y")
	}
}

func TestLineMove(test *testing.T) {
	line1 := Line{Point{100, 200}, Point{300, 400}}
	line2 := line1.Move(Vector{50, 60}).(Line)

	if line2.P1.X != 150 {
		test.Error("line move point 1 x")
	}
	if line2.P1.Y != 260 {
		test.Error("line move point 1 y")
	}
	if line2.P2.X != 350 {
		test.Error("line move point 2 x")
	}
	if line2.P2.Y != 460 {
		test.Error("line move point 2 y")
	}
}

func TestLineVector(test *testing.T) {
	line := Line{Point{450, 675}, Point{525, 700}}
	vector := line.Vector()

	if vector.DX != 75 {
		test.Error("test line vector dx")
	}
	if vector.DY != 25 {
		test.Error("test line vector dy")
	}
}

func TestLineEquation(test *testing.T) {
	line1 := Line{Point{130, 140}, Point{250, 140}}
	equation1 := line1.Equation()

	if equation1(200) != 140 {
		test.Error("Equation 1 invalid output: ", equation1(200))
	}

	line2 := Line{Point{165, 165}, Point{165, 200}}
	equation2 := line2.Equation()

	if equation2 != nil {
		test.Error("Equation 2 exists")
	}

	line3 := Line{Point{451, 363}, Point{149, 450}}
	equation3 := line3.Equation()

	if equation3(line3.P1.X) != line3.P1.Y {
		test.Error("Equation 3 does not satisfy p1")
	}
	if equation3(line3.P2.X) != line3.P2.Y {
		test.Error("Equation 3 does not satisfy p2")
	}
}

func TestYIntersect(test *testing.T) {
	line1 := Line{Point{0, 100}, Point{220, 450}}
	slope1 := line1.Vector().AsSlope()

	if YIntersect(slope1, line1.P1) != 100 {
		test.Error("test y intersect l1 p1")
	}
	if YIntersect(slope1, line1.P2) != 100 {
		test.Error("test y intersect l1 p2")
	}

	line2 := Line{Point{100, 250}, Point{200, 250}}
	slope2 := line2.Vector().AsSlope()

	if YIntersect(slope2, line2.P1) != 250 {
		test.Error("test y intersect l2 p1")
	}
	if YIntersect(slope2, line2.P2) != 250 {
		test.Error("test y intersect l2 p2")
	}
}
