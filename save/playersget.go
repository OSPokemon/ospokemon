package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	// "github.com/ospokemon/ospokemon/engine"
	// "github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
	// "time"
)

const EVNT_PlayersGet = "save.PlayersGet"

func PlayersGet(username string) (*Player, error) {
	p := PlayersNew(username)

	row := Connection.QueryRow(
		"SELECT level, experience, money FROM players WHERE username=?",
		p.Username,
	)

	if err := row.Scan(&p.Level, &p.Experience, &p.Money); err != nil {
		return nil, errors.New("playergetscan: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Debug("save.PlayersGet")

	util.Event.Fire(EVNT_PlayersGet, p)

	return p, nil
}
