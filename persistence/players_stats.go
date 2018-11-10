package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func PlayersStatsSelect(player *ospokemon.Player) (ospokemon.Stats, error) {
	rows, err := Connection.Query(
		"SELECT stat, value, base, base FROM players_stats WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, err
	}

	stats := make(ospokemon.Stats)

	for rows.Next() {
		var namebuff string
		stat := &ospokemon.Stat{}

		if err = rows.Scan(&namebuff, &stat.Value, &stat.Max, &stat.Base); err != nil {
			return stats, err
		}

		stats[namebuff] = stat
	}
	rows.Close()

	log.Add("Username", player.Username).Add("Stats", stats).Debug("players_stats select")

	return stats, nil
}

func PlayersStatsInsert(player *ospokemon.Player) error {
	stats := player.GetStats()
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

func PlayersStatsDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM players_stats WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Debug("players_stats delete")
	}

	return err
}
