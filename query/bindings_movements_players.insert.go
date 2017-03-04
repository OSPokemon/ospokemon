package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsMovementsPlayersInsert(player *game.Player, movements map[string]string) error {
	for key, direction := range movements {
		_, err := Connection.Exec(
			"INSERT INTO bindings_movements_players (username, key, direction) VALUES (?, ?, ?)",
			player.Username,
			key,
			direction,
		)

		if err != nil {
			return err
		}
	}

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Menus":    movements,
	}).Debug("bindings_movements_players insert")

	return nil
}
