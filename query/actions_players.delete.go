package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func ActionsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", player.Username)

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
	}).Debug("actions_players delete")

	return err
}
