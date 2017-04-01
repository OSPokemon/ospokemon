package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
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
