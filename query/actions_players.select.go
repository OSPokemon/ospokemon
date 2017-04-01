package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"time"
)

func ActionsPlayersSelect(player *ospokemon.Player) (ospokemon.Actions, error) {
	rows, err := Connection.Query(
		"SELECT spell, timer FROM actions_players WHERE username=?",
		player.Username,
	)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	actions := make(ospokemon.Actions)

	for rows.Next() {
		var spellbuff uint
		var timebuff uint64
		err = rows.Scan(&spellbuff, &timebuff)

		if err != nil {
			return nil, err
		}

		if spell, err := GetSpell(spellbuff); spell != nil {
			action := ospokemon.BuildAction(spell)

			if t := time.Duration(timebuff); timebuff > 0 {
				action.Timer = &t
			} else {
				action.Timer = nil
			}

			actions[action.Spell.Id] = action
		} else {
			return nil, err
		}
	}

	log.Add("Username", player.Username).Add("Actions", actions).Debug("actions_players select")

	event.Fire(event.ActionsPlayersSelect, player, actions)

	return actions, nil
}
