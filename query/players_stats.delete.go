package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

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
