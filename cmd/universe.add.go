package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/comp"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(engine.EVNT_UniverseAdd, UniverseAdd)
}

func UniverseAdd(args ...interface{}) {
	e := args[0].(*engine.Entity)
	l := e.Component(comp.LOCATION).(*comp.Location)
	log := logrus.WithFields(logrus.Fields{
		"Universe": l.UniverseId,
		"Entity":   e.Id,
	})

	if err := universeadd(e, l); err != nil {
		log.Error("cmd.UniverseAdd: " + err.Error())
	} else {
		log.Info("cmd.UniverseAdd")
	}
}

func universeadd(e *engine.Entity, l *comp.Location) error {

	if engine.Multiverse[l.UniverseId] == nil {
		if u, err := save.UniversesGet(l.UniverseId); err == nil {
			engine.Multiverse[u.Id] = u
		} else {
			return err
		}
	}

	u := engine.Multiverse[l.UniverseId]
	e.Id = u.GenerateId()
	u.Entities[e.Id] = e

	return nil
}
