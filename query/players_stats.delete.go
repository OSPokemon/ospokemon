package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersStatsDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM players_stats WHERE username=?",
		player.Username,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players_stats delete")
	}

	return err
}
