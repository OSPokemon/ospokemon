package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
)

const EVNT_PlayersDelete = "save.PlayersDelete"

func PlayersDelete(username string) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", username)

	if err != nil {
		return errors.New("playersdelete: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Debug("save.PlayersDelete")

	util.Event.Fire(EVNT_PlayersDelete, username)

	return nil
}
