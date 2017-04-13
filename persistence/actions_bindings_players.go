package persistence

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func ActionsBindingsPlayersSelect(player *ospokemon.Player) (map[string]uint, error) {
	rows, err := Connection.Query(
		"SELECT spell, key FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return nil, err
	}

	bindings := make(map[string]uint)

	for rows.Next() {
		var spellbuff uint
		var keybuff string

		if err = rows.Scan(&spellbuff, &keybuff); err != nil {
			return bindings, err
		}

		bindings[keybuff] = spellbuff
	}
	rows.Close()

	log.Add("Username", player.Username).Add("Bindings", bindings).Debug("actions_bindings_players select")

	return bindings, nil
}

func ActionsBindingsPlayersInsert(player *ospokemon.Player, insert map[string]uint) error {
	for key, action := range insert {
		_, err := Connection.Exec(
			"INSERT INTO actions_bindings_players (username, key, spell) VALUES (?, ?, ?)",
			player.Username,
			key,
			action,
		)

		if err != nil {
			return err
		}
	}

	log.Add("Username", player.Username).Add("Bindings", insert).Debug("actions_bindings_players insert")

	return nil
}

func ActionsBindingsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("actions_bindings_players delete")
	}

	return err
}