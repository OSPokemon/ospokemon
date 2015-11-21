package world

func CreatePathVector(pos1 Position, pos2 Position, speed int) *Vector {
	// do some pathing
	return &Vector{
		Radians:  0,
		Distance: float64(speed) / 1000,
	}
}
