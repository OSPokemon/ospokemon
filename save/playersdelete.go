package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func PlayersDelete(username string) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", username)

	if err != nil {
		return errors.New("playersdelete: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Debug("save.PlayersDelete")

	return nil
}
