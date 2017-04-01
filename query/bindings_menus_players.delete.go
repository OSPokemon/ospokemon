package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsMenusPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_menus_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Debug("bindings_menus_players delete")
	}

	return err
}
