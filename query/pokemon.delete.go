package query

import (
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func PokemonDelete(id uint) error {
	_, err := Connection.Exec("DELETE FROM pokemon WHERE id=?", id)

	if err == nil {
		log.Add("Pokemon", id).Info("pokemon delete")

		event.Fire(event.PokemonDelete, id)
	}

	return err
}
