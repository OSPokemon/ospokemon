package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func PlayersInsert(player *game.Player) error {
	_, err := Connection.Exec(
		"INSERT INTO players (username, level, experience, money, class, bagsize) values (?, ?, ?, ?, ?, ?)",
		player.Username,
		player.Level,
		player.Experience,
		player.Money,
		player.Class,
		player.BagSize,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Info("players insert")

		delete(game.Players, player.Username)
		event.Fire(event.PlayersInsert, player)
	}

	return err
}
