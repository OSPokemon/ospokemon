package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
)

func PokemonDelete(id uint) error {
	_, err := Connection.Exec("DELETE FROM pokemon WHERE id=?", id)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Pokemon": id,
		}).Info("pokemon delete")

		event.Fire(event.PokemonDelete, id)
	}

	return err
}
