package persistence

import (
	"github.com/ospokemon/ospokemon"
)

func BindingsMenusPlayersSelect(player *ospokemon.Player) (map[string]string, error) {
	rows, err := Connection.Query(
		"SELECT `key`, menu FROM bindings_menus_players WHERE username=?",
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

	ospokemon.LOG().Add("Username", player.Username).Add("Bindings", menus).Debug("bindings_menus_players select")

	return menus, nil
}

func BindingsMenusPlayersInsert(player *ospokemon.Player) error {
	menus := make(map[string]ospokemon.Menu)

	if bindings := player.GetBindings(); bindings != nil {
		for key, binding := range bindings {
			if menu := binding.GetMenu(); menu != "" {
				menus[key] = menu
			}
		}
	}

	for key, menu := range menus {
		_, err := Connection.Exec(
			"INSERT INTO bindings_menus_players (username, `key`, menu) VALUES (?, ?, ?)",
			player.Username,
			key,
			string(menu),
		)

		if err != nil {
			return err
		}
	}

	ospokemon.LOG().Add("Username", player.Username).Add("Menus", menus).Debug("bindings_menus_players insert")

	return nil
}

func BindingsMenusPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_menus_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		ospokemon.LOG().Add("Username", player.Username).Debug("bindings_menus_players delete")
	}

	return err
}
