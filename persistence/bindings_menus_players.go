package persistence

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsMenusPlayersSelect(player *ospokemon.Player) (map[string]string, error) {
	rows, err := Connection.Query(
		"SELECT key, menu FROM bindings_menus_players WHERE username=?",
		player.Username,
	)

	menus := make(map[string]string)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var keybuff, menubuff string

		if err := rows.Scan(&keybuff, &menubuff); err != nil {
			return menus, err
		}

		menus[keybuff] = menubuff
	}
	rows.Close()

	log.Add("Username", player.Username).Add("Bindings", menus).Debug("bindings_menus_players select")

	return menus, nil
}

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
