package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsMenusPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_menus_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("bindings_menus_players delete")
	}

	return err
}
