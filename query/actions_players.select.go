package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"time"
)

func ActionsPlayersSelect(player *game.Player) (game.Actions, error) {
	rows, err := Connection.Query(
		"SELECT spell, timer FROM actions_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return nil, err
	}

	actions := make(game.Actions)

	for rows.Next() {
		action := game.MakeAction()
		var timebuff uint64
		err = rows.Scan(&action.Spell, &timebuff)

		if err != nil {
			return actions, err
		}

		if t := time.Duration(timebuff); timebuff > 0 {
			action.Timer = &t
		} else {
			action.Timer = nil
		}

		event.Fire(event.ActionBuild, action)

		actions[action.Spell] = action
	}
	rows.Close()

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Actions":  actions,
	}).Debug("actions_players select")

	event.Fire(event.ActionsPlayersSelect, player, actions)
	return actions, nil
}
