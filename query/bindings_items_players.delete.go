package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func BindingsItemsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("bindings_items_players delete")
	}

	return err
}
