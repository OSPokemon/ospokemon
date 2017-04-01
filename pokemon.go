package ospokemon

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

func MakePokemon() *Pokemon {
	p := &Pokemon{
		Parts: make(Parts),
	}

	return p
}
