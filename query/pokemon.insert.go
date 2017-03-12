package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func PokemonInsert(pokemon *game.Pokemon) error {
	_, err := Connection.Exec(
		"INSERT INTO pokemon (id, species, name, xp, level, gender, shiny) VALUES (?, ?, ?, ?, ?, ?, ?)",
		pokemon.Id,
		pokemon.Species,
		pokemon.Name,
		pokemon.Xp,
		pokemon.Level,
		pokemon.Gender,
		pokemon.Shiny,
	)

	if err == nil {
		log.Add("Pokemon", pokemon.Id).Info("pokemon insert")

		event.Fire(event.PokemonInsert, pokemon)
	}

	return err
}
