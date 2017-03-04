package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func BindingsMovementsPlayersSelect(player *game.Player) (map[string]string, error) {
	rows, err := Connection.Query(
		"SELECT key, direction FROM bindings_movements_players WHERE username=?",
		player.Username,
	)

	movements := make(map[string]string)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var keybuff, directionbuff string

		if err := rows.Scan(&keybuff, &directionbuff); err != nil {
			return movements, err
		}

		movements[keybuff] = directionbuff
	}
	rows.Close()

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Bindings": movements,
	}).Debug("bindings_movements_players select")

	return movements, nil
}
