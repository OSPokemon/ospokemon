package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func PlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", player.Username)

	if err == nil {
		log.Add("Username", player.Username).Info("players delete")

		event.Fire(event.PlayersDelete, player)
	}

	return err
}
