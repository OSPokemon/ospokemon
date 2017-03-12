package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func PlayersStatsDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM players_stats WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", "2").Debug("players_stats delete")
	}

	return err
}
