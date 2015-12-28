package physics

import (
	"math"
	"testing"
)

func TestDistancePointPoint(test *testing.T) {
	point1 := Point{100, 100}
	point2 := Point{100, 200}

	if DistancePointPoint(point1, point2) != 100 {
		test.Error("Expected {100,100}<->{100,200} distance to be 100, got ", DistancePointPoint(point1, point2))
	}

	point3 := Point{200, 100}

	if DistancePointPoint(point1, point3) != 100 {
		test.Error("Expected {100,100}<->{200,100} distance to be 100, got ", DistancePointPoint(point1, point3))
	}

	point4 := Point{200, 200}

	if math.Floor(DistancePointPoint(point1, point4)) != 141 {
		test.Error("Expected {100,100}<->{200,200} distance to be 100, got ", DistancePointPoint(point1, point4))
	}

	point5 := Point{400, 500}

	if DistancePointPoint(point1, point5) != 500 {
		test.Error("Expected {100,100}<->{400,500} distance to be 500, got ", DistancePointPoint(point1, point5))
	}
}

func TestDistancePointLine(test *testing.T) {
	line1 := Line{Point{100, 100}, Point{100, 300}}
	point1 := Point{200, 200}
	if DistancePointLine(point1, line1) != 100 {
		test.Error("Distance to vertical line expected 100 got", DistancePointLine(point1, line1))
	}

	line2 := Line{Point{100, 100}, Point{300, 300}}
	point2 := Point{200, 200}
	if DistancePointLine(point2, line2) != 0 {
		test.Error("Distance to point on line expected 0 got", DistancePointLine(point2, line2))
	}

	line3 := Line{Point{100, 100}, Point{300, 100}} // horizontal
	point3 := Point{200, 400}
	if DistancePointLine(point3, line3) != 300 {
		test.Error("Distance to horizontal line expected 300 got", DistancePointLine(point3, line3))
	}

	line4 := Line{Point{100, 100}, Point{100, 300}}
	point4 := Point{200, 0}
	if math.Floor(DistancePointLine(point4, line4)) != 141 {
		test.Error("Distance to vertical edge expected 141 got", math.Floor(DistancePointLine(point4, line4)))
	}

	line5 := Line{Point{100, 100}, Point{200, 200}}
	point5 := Point{0, 0}
	if math.Floor(DistancePointLine(point5, line5)) != 141 {
		test.Error("Distance to line segment edge expected 141 got", math.Floor(DistancePointLine(point5, line5)))
	}

	line6 := Line{Point{100, 100}, Point{300, 100}}
	point6 := Point{0, 80}
	if math.Floor(DistancePointLine(point6, line6)) != 101 {
		test.Error("Distance to horizontal edge expected 101 got", math.Floor(DistancePointLine(point6, line6)))
	}

	line7 := Line{Point{0, 0}, Point{500, 500}}
	point7 := Point{350, 300}
	if math.Floor(DistancePointLine(point7, line7)) != 35 {
		test.Error("Distance close to line expected 35 got", math.Floor(DistancePointLine(point7, line7)))
	}
}

func TestDistanceLineLine(test *testing.T) {
	line1 := Line{Point{100, 100}, Point{500, 500}}
	line2 := Line{Point{100, 500}, Point{500, 100}}
	if distance := DistanceLineLine(line1, line2); distance != 0 {
		test.Error("Distance between crossing lines expected 0 got", distance)
	}

	line3 := Line{Point{100, 100}, Point{200, 200}}
	line4 := Line{Point{100, 0}, Point{200, 100}}
	if distance := math.Floor(DistanceLineLine(line3, line4)); distance != 70 {
		test.Error("Distance between parallel lines expected 70 got", distance)
	}

	line5 := Line{Point{100, 100}, Point{100, 300}}
	line6 := Line{Point{0, 200}, Point{200, 200}}
	if distance := DistanceLineLine(line5, line6); distance != 0 {
		test.Error("Distance between vertical/horizontal crossing lines expected 0 got", distance)
	}

	line7 := Line{Point{100, 100}, Point{100, 500}}
	line8 := Line{Point{0, 0}, Point{500, 0}}

	if distance := DistanceLineLine(line7, line8); distance != 100 {
		test.Error("Distance between vertical/horizontal lines expected 100 got", distance)
	}
}
