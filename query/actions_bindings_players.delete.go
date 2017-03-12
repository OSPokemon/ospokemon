package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func ActionsBindingsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("actions_bindings_players delete")
	}

	return err
}
