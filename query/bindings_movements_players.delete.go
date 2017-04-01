package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func BindingsMovementsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_movements_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Debug("bindings_movements_players delete")
	}

	return err
}
