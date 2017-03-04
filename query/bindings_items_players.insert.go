package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsItemsPlayersInsert(player *game.Player, items map[string]uint) error {
	for key, item := range items {
		_, err := Connection.Exec(
			"INSERT INTO bindings_items_players (username, key, item) VALUES (?, ?, ?)",
			player.Username,
			key,
			item,
		)

		if err != nil {
			return err
		}
	}

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Items":    items,
	}).Debug("bindings_items_players insert")

	return nil
}
