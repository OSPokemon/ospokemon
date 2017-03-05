package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsItemsPlayersInsert(player *game.Player, itemslots map[string]int) error {
	for key, itemslotid := range itemslots {
		_, err := Connection.Exec(
			"INSERT INTO bindings_items_players (username, key, itemslot) VALUES (?, ?, ?)",
			player.Username,
			key,
			itemslotid,
		)

		if err != nil {
			return err
		}
	}

	logrus.WithFields(logrus.Fields{
		"Username":  player.Username,
		"Itemslots": itemslots,
	}).Debug("bindings_items_players insert")

	return nil
}
