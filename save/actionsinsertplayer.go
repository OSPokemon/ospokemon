package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func ActionsInsertPlayer(username string, a *Action) error {
	timebuff := 0
	if a.Timer != nil {
		timebuff = int(*a.Timer)
	}

	_, err := Connection.Exec(
		"INSERT INTO actions_players (username, spellid, timer) VALUES (?, ?, ?)",
		username,
		a.SpellId,
		timebuff,
	)

	if err != nil {
		return errors.New("actionsinsertplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Actions":  a,
	}).Debug("save.ActionsInsertPlayer")

	return nil
}
