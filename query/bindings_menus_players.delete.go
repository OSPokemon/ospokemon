package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func BindingsMenusPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_menus_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Debug("bindings_menus_players delete")
	}

	return err
}
