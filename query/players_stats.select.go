package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersStatsSelect(player *game.Player) (game.Stats, error) {
	rows, err := Connection.Query(
		"SELECT stat, value, base, base FROM players_stats WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, err
	}

	stats := make(game.Stats)

	for rows.Next() {
		var namebuff string
		stat := &game.Stat{}

		if err = rows.Scan(&namebuff, &stat.Value, &stat.Max, &stat.Base); err != nil {
			return stats, err
		}

		stats[namebuff] = stat
	}
	rows.Close()

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Stats":    stats,
	}).Debug("players_stats select")

	return stats, nil
}
