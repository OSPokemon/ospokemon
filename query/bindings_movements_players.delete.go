package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
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
