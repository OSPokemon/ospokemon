package registry

import (
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/world"
)

var Players = make(map[int]*entities.Player)

var Pokemon = make(map[int]*entities.PokemonEntity)

var Scripts = make(map[string]world.SpellScript)
