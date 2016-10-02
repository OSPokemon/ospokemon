package spell

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	engine.Scripts[0] = Super
}

func Super(u *engine.Universe, e *engine.Entity, data map[string]string) {
	p := e.Component(save.COMP_Player).(*save.Player)

	_, err := save.Connection.Exec("DELETE FROM actions_players WHERE username=?",
		p.Username,
	)

	if err != nil {
		logrus.Error(err)
		return
	}

	_, err = save.Connection.Exec("INSERT INTO actions_players (username, spellid, timer) SELECT ?, id, 0 FROM spells",
		p.Username,
	)

	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"Universe": u.Id,
		"Entity":   e.Id,
	}).Warn("script.Super")

	util.Event.Fire(save.EVNT_PlayerPull, p.Username)
}
