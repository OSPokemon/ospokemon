package space

type Shape interface {
	Center() Point
	Copy() Shape
	Move(Vector) Shape
}
