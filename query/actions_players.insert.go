package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func ActionsPlayersInsert(player *game.Player, actions game.Actions) error {
	for spell, action := range actions {
		timebuff := 0
		if action.Timer != nil {
			timebuff = int(*action.Timer)
		}

		_, err := Connection.Exec(
			"INSERT INTO actions_players (username, spell, timer) VALUES (?, ?, ?)",
			player.Username,
			spell,
			timebuff,
		)

		if err != nil {
			return err
		}
	}

	log.Add("Username", player.Username).Add("Actions", actions).Debug("actions_players insert")

	event.Fire(event.ActionsPlayersInsert, player, actions)
	return nil
}
