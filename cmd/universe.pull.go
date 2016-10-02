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
	log := logrus.WithFields(logrus.Fields{
		"UniverseId": universeid,
	})

	if err := universepull(universeid); err != nil {
		log.Error("cmd.UniversePull: " + err.Error())
	} else {
		log.Info("cmd.UniversePull")
	}
}

func universepull(universeid uint) error {
	u, err := universepullall(universeid)
	engine.Multiverse[universeid] = u
	return err
}

func universepullall(universeid uint) (*engine.Universe, error) {
	u := engine.MakeUniverse(universeid)

	if err := universepullbase(u); err != nil {
		return nil, err
	}

	if err := universepulllayout(u); err != nil {
		return nil, err
	}

	return u, nil
}

func universepullbase(u *engine.Universe) error {
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
