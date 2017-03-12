package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func BindingsItemsPlayersSelect(player *game.Player) (map[string]int, error) {
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
