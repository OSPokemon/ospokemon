package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
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

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Menus":    menus,
	}).Debug("bindings_menus_players insert")

	return nil
}
