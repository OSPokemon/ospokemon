package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsMovementsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_movements_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("bindings_movements_players delete")
	}

	return err
}
