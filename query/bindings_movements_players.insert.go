package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func BindingsMovementsPlayersInsert(player *game.Player, movements map[string]string) error {
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
