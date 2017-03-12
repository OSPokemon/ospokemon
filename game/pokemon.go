package game

const PARTpokemon = "pokemon"

type Pokemon struct {
	Id      uint
	Species uint
	Name    string
	Xp      uint
	Level   uint
	Gender  string
	Shiny   bool
	Parts
}

var Pokemons = make(map[uint]*Pokemon)

func MakePokemon(id uint) *Pokemon {
	p := &Pokemon{
		Id:    id,
		Parts: make(Parts),
	}

	return p
}
