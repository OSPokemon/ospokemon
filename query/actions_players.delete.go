package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func ActionsPlayersDelete(player *game.Player) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", player.Username)

	log.Add("Username", player.Username).Debug("actions_players delete")

	return err
}
