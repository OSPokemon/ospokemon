package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/log"
)

func PokemonDelete(id uint) error {
	_, err := Connection.Exec("DELETE FROM pokemon WHERE id=?", id)

	if err == nil {
		log.Add("Pokemon", "2").Info("pokemon delete")

		event.Fire(event.PokemonDelete, id)
	}

	return err
}
