package game

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/part"
)

type Pokemon struct {
	Id      uint
	Species uint
	Name    string
	Xp      uint
	Level   uint
	Gender  string
	Shiny   bool
	part.Parts
}

var Pokemons = make(map[uint]*Pokemon)

func MakePokemon(id uint) *Pokemon {
	p := &Pokemon{
		Id:    id,
		Parts: make(part.Parts),
	}

	logrus.Debug("Pokemon created")

	return p
}
