package persistence

import (
	"github.com/ospokemon/ospokemon"
)

func ActionsBindingsPlayersSelect(player *ospokemon.Player) (map[string]uint, error) {
	rows, err := Connection.Query(
		"SELECT spell, `key` FROM actions_bindings_players WHERE username=?",
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

	ospokemon.LOG().Add("Username", player.Username).Add("Bindings", bindings).Debug("actions_bindings_players select")

	return bindings, nil
}

func ActionsBindingsPlayersInsert(player *ospokemon.Player) error {
	actions := make(map[string]uint)

	if bindings := player.GetBindings(); bindings != nil {
		for key, binding := range bindings {
			if action := binding.GetAction(); action != nil {
				actions[key] = action.Spell.Id
			}
		}
	}

	for key, action := range actions {
		_, err := Connection.Exec(
			"INSERT INTO actions_bindings_players (username, `key`, spell) VALUES (?, ?, ?)",
			player.Username,
			key,
			action,
		)

		if err != nil {
			return err
		}
	}

	ospokemon.LOG().Add("Username", player.Username).Add("Bindings", actions).Debug("actions_bindings_players insert")

	return nil
}

func ActionsBindingsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM actions_bindings_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		ospokemon.LOG().Add("Username", player.Username).Debug("actions_bindings_players delete")
	}

	return err
}
