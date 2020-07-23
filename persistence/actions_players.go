package persistence

import (
	"time"

	"github.com/ospokemon/ospokemon"
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

		if spell, err := ospokemon.GetSpell(spellbuff); spell != nil {
			action := ospokemon.BuildAction(spell)

			if t := time.Duration(timebuff); timebuff > 0 {
				action.Timer.Set(t)
			} else {
				action.Timer = nil
			}

			actions[action.Spell.Id] = action
		} else {
			return nil, err
		}
	}

	ospokemon.LOG().Add("Username", player.Username).Add("Actions", actions).Debug("actions_players select")

	return actions, nil
}

func ActionsPlayersInsert(player *ospokemon.Player) error {
	actions := player.GetActions()

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

	ospokemon.LOG().Add("Username", player.Username).Add("Actions", actions).Debug("actions_players insert")
	return nil
}

func ActionsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", player.Username)

	ospokemon.LOG().Add("Username", player.Username).Debug("actions_players delete")

	return err
}
