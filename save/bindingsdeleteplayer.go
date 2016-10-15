package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func BindingsDeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM bindings_players WHERE username=?", username)

	if err != nil {
		return errors.New("bindingsdeleteplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Debug("save.BindingsDeletePlayer")

	return nil
}
