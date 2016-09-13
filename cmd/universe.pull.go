package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(engine.EVNT_UniversePull, UniversePull)
}

func UniversePull(args ...interface{}) {
	universeid := args[0].(uint)

	u, err := universepullall(universeid)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"UniverseId": universeid,
		}).Error("cmd.UniversePull: " + err.Error())
		return
	}

	engine.Multiverse[universeid] = u

	logrus.WithFields(logrus.Fields{
		"UniverseId": universeid,
	}).Info("cmd.UniversePull")
}

func universepullall(universeid uint) (*engine.Universe, error) {
	u := engine.MakeUniverse(universeid)

	if err := universepull(u); err != nil {
		return nil, err
	}

	if err := universepulllayout(u); err != nil {
		return nil, err
	}

	return u, nil
}

func universepull(u *engine.Universe) error {
	row := save.Connection.QueryRow(
		"SELECT x, y, dx, dy FROM universes WHERE id=?",
		u.Id,
	)

	if err := row.Scan(&u.Space.Rect.Anchor.X, &u.Space.Rect.Anchor.Y, &u.Space.Rect.Dimension.DX, &u.Space.Rect.Dimension.DY); err != nil {
		return errors.New("universepull: " + err.Error())
	}

	return nil
}

func universepulllayout(u *engine.Universe) error {
	return nil
}
