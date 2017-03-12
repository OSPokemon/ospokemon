package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func BindingsMenusPlayersInsert(player *game.Player, menus map[string]string) error {
	for key, menu := range menus {
		_, err := Connection.Exec(
			"INSERT INTO bindings_menus_players (username, key, menu) VALUES (?, ?, ?)",
			player.Username,
			key,
			menu,
		)

		if err != nil {
			return err
		}
	}

	log.Add("Username", player.Username).Add("Menus", menus).Debug("bindings_menus_players insert")

	return nil
}
