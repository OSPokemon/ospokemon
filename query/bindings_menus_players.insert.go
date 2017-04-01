package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsMenusPlayersInsert(player *ospokemon.Player, menus map[string]string) error {
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
