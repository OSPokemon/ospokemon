package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func PlayersStatsInsert(player *ospokemon.Player, stats ospokemon.Stats) error {
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

	log.Add("Username", player.Username).Add("stats", stats).Debug("players_stats insert")

	return nil
}
