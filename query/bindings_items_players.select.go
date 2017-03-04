package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsItemsPlayersSelect(player *game.Player) (map[string]uint, error) {
	rows, err := Connection.Query(
		"SELECT item, key FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return nil, err
	}

	items := make(map[string]uint)

	for rows.Next() {
		var itembuff uint
		var keybuff string

		if err = rows.Scan(&itembuff, &keybuff); err != nil {
			return items, err
		}

		items[keybuff] = itembuff
	}
	rows.Close()

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Items":    items,
	}).Debug("bindings_items_players select")

	return items, nil
}
