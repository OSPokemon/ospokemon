package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func ActionsBindingsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("actions_bindings_players delete")
	}

	return err
}
