package query

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func ItembagsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM itemslots_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return err
	}

	_, err = Connection.Exec(
		"DELETE FROM itembags_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		log.Add("Username", player.Username).Debug("itembags_players delete")

		event.Fire(event.ItembagsPlayersDelete, player)
	}

	return err
}
