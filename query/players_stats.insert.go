package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersStatsInsert(player *game.Player, stats game.Stats) error {
	for name, stat := range stats {
		_, err := Connection.Exec(
			"INSERT INTO players_stats (username, stat, value, base) VALUES (?, ?, ?, ?)",
			player.Username,
			name,
			stat.Value,
			stat.Base,
		)

		if err != nil {
			return err
		}
	}

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"stats":    stats,
	}).Debug("players_stats insert")

	return nil
}
