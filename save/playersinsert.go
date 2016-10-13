package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
)

const EVNT_PlayersInsert = "save.PlayersInsert"

func PlayersInsert(p *Player) error {
	_, err := Connection.Exec(
		"INSERT INTO players (username, level, experience, money) values (?, ?, ?, ?)",
		p.Username,
		p.Level,
		p.Experience,
		p.Money,
	)

	util.Event.Fire(EVNT_PlayersInsert, p)

	if err != nil {
		return errors.New("playersinsert: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": p.Username,
	}).Debug("save.PlayersInsert")

	return nil
}
