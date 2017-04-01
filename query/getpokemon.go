package query

import (
	"ospokemon.com"
)

func GetPokemon(id uint) (*ospokemon.Pokemon, error) {
	if pokemon, ok := ospokemon.Pokemons[id]; ok {
		return pokemon, nil
	}

	return PokemonSelect(id)
}
