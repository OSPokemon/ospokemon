package persistence

import (
	"ospokemon.com"
	"ospokemon.com/log"
)

func BindingsItemsPlayersSelect(player *ospokemon.Player) (map[string]int, error) {
	rows, err := Connection.Query(
		"SELECT key, itemslot FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return nil, err
	}

	itemslots := make(map[string]int)

	for rows.Next() {
		var keybuff string
		var itemslot int

		if err = rows.Scan(&keybuff, &itemslot); err != nil {
			return itemslots, err
		}

		itemslots[keybuff] = itemslot
	}
	rows.Close()

	log.Add("Username", player.Username).Add("Itemslots", itemslots).Debug("bindings_items_players select")

	return itemslots, nil
}

func BindingsItemsPlayersInsert(player *ospokemon.Player, itemslots map[string]int) error {
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

	log.Add("Username", player.Username).Add("Itemslots", itemslots).Debug("bindings_items_players insert")

	return nil
}

func BindingsItemsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("bindings_items_players delete")
	}

	return err
}
