package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func PlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", player.Username)

	if err == nil {
		log.Add("Username", player.Username).Info("players delete")

		event.Fire(event.PlayersDelete, player)
	}

	return err
}
