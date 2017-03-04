package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func EntitiesPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM entities_players WHERE username=?",
		player.Username,
	)

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
	}).Debug("entities_players delete")

	return err
}
