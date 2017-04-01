package query

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
