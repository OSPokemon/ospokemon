package persistence

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsMovementsPlayersSelect(player *ospokemon.Player) (map[string]string, error) {
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

	log.Add("Username", player.Username).Add("Bindings", movements).Debug("bindings_movements_players select")

	return movements, nil
}

func BindingsMovementsPlayersInsert(player *ospokemon.Player, movements map[string]string) error {
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

	log.Add("Username", player.Username).Add("Menus", movements).Debug("bindings_movements_players insert")

	return nil
}

func BindingsMovementsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM bindings_movements_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Debug("bindings_movements_players delete")
	}

	return err
}
