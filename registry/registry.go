package registry

import (
	"github.com/ospokemon/ospokemon/objects/auth"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/world"
)

var Accounts = make(map[string]*auth.Account)

var Players = make(map[int]*entities.Player)

var Pokemon = make(map[int]*entities.PokemonEntity)

var Scripts = make(map[string]world.SpellScript)
