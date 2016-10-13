package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

func ActionsGetPlayer(username string) (map[uint]*Action, error) {
	a := make(map[uint]*Action)
	rows, err := Connection.Query(
		"SELECT spellid, timer FROM actions_players WHERE username=?",
		username,
	)

	if err != nil {
		return nil, err
	} else {
		defer rows.Close()
	}

	for rows.Next() {
		var timebuff uint64
		action := &Action{}

		err = rows.Scan(&action.SpellId, &timebuff)
		if err != nil {
			return nil, errors.New("actionsgetplayer: " + err.Error())
		}

		if t := time.Duration(timebuff); timebuff > 0 {
			action.Timer = &t
		}

		a[action.SpellId] = action
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Actions":  a,
	}).Debug("save.ActionsGetPlayer")

	return a, nil
}
