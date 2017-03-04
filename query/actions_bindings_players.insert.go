package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func ActionsBindingsPlayersInsert(player *game.Player, insert map[string]uint) error {
	for key, action := range insert {
		_, err := Connection.Exec(
			"INSERT INTO actions_bindings_players (username, key, spell) VALUES (?, ?, ?)",
			player.Username,
			key,
			action,
		)

		if err != nil {
			return err
		}
	}

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Bindings": insert,
	}).Debug("actions bindings players insert")

	return nil
}
