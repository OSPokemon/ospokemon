package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetPokemon(id uint) (*game.Pokemon, error) {
	if pokemon, ok := game.Pokemons[id]; ok {
		return pokemon, nil
	}

	return PokemonSelect(id)
}
