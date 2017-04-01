package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
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

	log.Add("Username", player.Username).Add("Itemslots", itemslots).Debug("bindings_items_players insert")

	return nil
}
