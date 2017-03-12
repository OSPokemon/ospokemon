package query

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func ActionsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", player.Username)

	log.Add("Username", player.Username).Debug("actions_players delete")

	return err
}
