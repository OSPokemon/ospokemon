package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(engine.EVNT_UniverseRemove, UniverseRemove)
}

func UniverseRemove(args ...interface{}) {
	e := args[0].(*engine.Entity)
	l := e.Component(engine.COMP_Location).(*engine.Location)
	log := logrus.WithFields(logrus.Fields{
		"Universe": l.UniverseId,
		"Entity":   e.Id,
	})

	if err := universeremove(e, l); err != nil {
		log.Error("cmd.UniverseRemove: " + err.Error())
	} else {
		log.Info("cmd.UniverseRemove")
	}
}

func universeremove(e *engine.Entity, l *engine.Location) error {
	u := engine.Multiverse[l.UniverseId]

	if u == nil {
		return errors.New("Universe not found")
	}

	delete(u.Entities, e.Id)
	e.Id = 0

	return nil
}
