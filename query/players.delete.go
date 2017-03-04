package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", player.Username)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Info("players delete")

		event.Fire(event.PlayersDelete, player)
	}

	return err
}
