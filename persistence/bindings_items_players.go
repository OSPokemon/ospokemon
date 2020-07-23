package persistence

import (
	"github.com/ospokemon/ospokemon"
)

func BindingsItemsPlayersSelect(player *ospokemon.Player) (map[string]uint, error) {
	rows, err := Connection.Query(
		"SELECT `key`, itemid FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return nil, err
	}

	itemslots := make(map[string]uint)

	for rows.Next() {
		var keybuff string
		var itemidbuf uint

		if err = rows.Scan(&keybuff, &itemidbuf); err != nil {
			return itemslots, err
		}

		itemslots[keybuff] = itemidbuf
	}
	rows.Close()

	ospokemon.LOG().Add("Username", player.Username).Add("Bindings", itemslots).Debug("bindings_items_players select")

	return itemslots, nil
}

func BindingsItemsPlayersInsert(player *ospokemon.Player) error {
	itemslots := make(map[string]uint)

	if bindings := player.GetBindings(); bindings != nil {
		for key, binding := range bindings {
			if itemslot := binding.GetItemslot(); itemslot != nil {
				itemslots[key] = itemslot.Item.Id
			}
		}
	}

	for key, itemid := range itemslots {
		_, err := Connection.Exec(
			"INSERT INTO bindings_items_players (username, `key`, itemid) VALUES (?, ?, ?)",
			player.Username,
			key,
			itemid,
		)

		if err != nil {
			return err
		}
	}

	ospokemon.LOG().Add("Username", player.Username).Add("Itemslots", itemslots).Debug("bindings_items_players insert")

	return nil
}

func BindingsItemsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?",
		player.Username,
	)

	if err == nil {
		ospokemon.LOG().Add("Username", player.Username).Debug("bindings_items_players delete")
	}

	return err
}
