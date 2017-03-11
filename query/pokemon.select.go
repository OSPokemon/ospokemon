package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func PokemonSelect(id uint) (*game.Pokemon, error) {
	row := Connection.QueryRow(
		"SELECT species, name, xp, level, gender, shiny FROM pokemon WHERE id=?",
		id,
	)

	pokemon := game.MakePokemon(id)

	err := row.Scan(
		&pokemon.Species,
		&pokemon.Name,
		&pokemon.Xp,
		&pokemon.Level,
		&pokemon.Gender,
		&pokemon.Shiny,
	)

	if err != nil {
		game.Pokemons[id] = nil
		return nil, err
	}

	game.Pokemons[id] = pokemon

	logrus.WithFields(logrus.Fields{
		"Pokemon": id,
	}).Info("pokemon select")

	event.Fire(event.PokemonSelect, pokemon)
	return pokemon, nil
}