package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func BindingsGetPlayer(username string) (map[string]*Binding, error) {
	bindings := make(map[string]*Binding)
	rows, err := Connection.Query(
		"SELECT key, spellid FROM bindings_players WHERE username=?",
		username,
	)

	if err != nil {
		return nil, errors.New("bindingsgetplayer: " + err.Error())
	} else {
		defer rows.Close()
	}

	for rows.Next() {
		binding := &Binding{}
		err = rows.Scan(&binding.Key, &binding.SpellId)

		if err != nil {
			return nil, errors.New("bindingsgetplayer: " + err.Error())
		}

		bindings[binding.Key] = binding
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Bindings": bindings,
	}).Debug("save.BindingsGetPlayer")

	return bindings, nil
}
