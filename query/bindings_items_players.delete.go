package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsItemsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Info("bindings_items_players delete")
	}

	return err
}
