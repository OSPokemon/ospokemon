package space

import (
	"math"
)

type Vector struct {
	DX float64
	DY float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		DX: v1.DX + v2.DX,
		DY: v1.DY + v2.DY,
	}
}

func (v Vector) Copy() Vector {
	return Vector{
		DX: v.DX,
		DY: v.DY,
	}
}

func (v Vector) Multiply(f float64) Vector {
	return Vector{
		DY: v.DY * f,
		DX: v.DX * f,
	}
}

func (v Vector) Reverse() Vector {
	return Vector{
		DY: -v.DY,
		DX: -v.DX,
	}
}

func (v Vector) MakeUnit() Vector {
	length := math.Sqrt((v.DX * v.DX) + (v.DY * v.DY))

	if math.IsNaN(length) || length == 0 {
		length = 1
	}

	return Vector{
		DY: v.DY / length,
		DX: v.DX / length,
	}
}

func (v Vector) AsSlope() float64 {
	if v.DX == 0 {
		return math.Inf(1)
	} else {
		return v.DY / v.DX
	}
}

func (v Vector) MakeRightAngle() Vector {
	return Vector{
		DY: v.DX,
		DX: -v.DY,
	}
}
