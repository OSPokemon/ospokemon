package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func PokemonSelect(id uint) (*ospokemon.Pokemon, error) {
	row := Connection.QueryRow(
		"SELECT id, species, name, xp, level, gender, shiny FROM pokemon WHERE id=?",
		id,
	)

	pokemon := ospokemon.MakePokemon()

	err := row.Scan(
		&pokemon.Id,
		&pokemon.Species,
		&pokemon.Name,
		&pokemon.Xp,
		&pokemon.Level,
		&pokemon.Gender,
		&pokemon.Shiny,
	)

	if err != nil {
		ospokemon.Pokemons[id] = nil
		return nil, err
	}

	ospokemon.Pokemons[id] = pokemon

	log.Add("Pokemon", id).Info("pokemon select")

	event.Fire(event.PokemonSelect, pokemon)
	return pokemon, nil
}
