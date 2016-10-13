package save

import (
	"github.com/ospokemon/ospokemon/engine"
)

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Entity     *engine.Entity
}

var Players = make(map[string]*Player)
