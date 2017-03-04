package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsMenusPlayersSelect(player *game.Player) (map[string]string, error) {
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

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Bindings": menus,
	}).Debug("bindings_menus_players select")

	return menus, nil
}
