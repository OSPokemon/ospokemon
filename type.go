package ospokemon

type Type struct {
	Id     uint
	Image  string
	Strong []uint
}

var Types = make(map[uint]*Type)

func MakeType(id uint) *Type {
	return &Type{
		Id:     id,
		Strong: make([]uint, 0),
	}
}
