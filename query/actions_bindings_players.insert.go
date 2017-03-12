package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func ActionsBindingsPlayersInsert(player *game.Player, insert map[string]uint) error {
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
