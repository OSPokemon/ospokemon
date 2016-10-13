package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func BindingsInsertPlayer(username string, b *Binding) error {
	_, err := Connection.Exec(
		"INSERT INTO bindings_players (username, key, spellid) VALUES (?, ?, ?)",
		username,
		b.Key,
		b.SpellId,
	)

	if err != nil {
		return errors.New("locationsinsertplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Binding":  b,
	}).Debug("save.BindingsInsertPlayer")

	return nil
}
