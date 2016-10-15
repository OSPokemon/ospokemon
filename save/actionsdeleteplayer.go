package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func ActionsDeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", username)

	if err != nil {
		return errors.New("actionsdeleteplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Debug("save.ActionsDeletePlayer")

	return nil
}
