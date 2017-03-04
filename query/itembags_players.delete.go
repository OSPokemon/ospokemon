package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func ItembagsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM itemslots_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return err
	}

	_, err = Connection.Exec(
		"DELETE FROM itembags_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("itembags_players delete")

		event.Fire(event.ItembagsPlayersDelete, player)
	}

	return err
}
