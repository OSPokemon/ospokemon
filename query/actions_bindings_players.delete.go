package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func ActionsBindingsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Info("actions bindings players delete")
	}

	return err
}
