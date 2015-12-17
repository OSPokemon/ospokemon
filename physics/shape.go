package physics

type Shape interface {
	Copy() Shape
	Move(Vector) Shape
}
