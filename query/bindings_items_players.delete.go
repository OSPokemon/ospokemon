package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsItemsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("bindings_items_players delete")
	}

	return err
}
