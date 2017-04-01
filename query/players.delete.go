package query

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func PlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", player.Username)

	if err == nil {
		log.Add("Username", player.Username).Info("players delete")

		event.Fire(event.PlayersDelete, player)
	}

	return err
}
