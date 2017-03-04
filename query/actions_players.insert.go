package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
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

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Actions":  actions,
	}).Debug("actions_players insert")

	event.Fire(event.ActionsPlayersInsert, player, actions)
	return nil
}
