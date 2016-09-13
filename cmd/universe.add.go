package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(engine.EVNT_UniverseAdd, UniverseAdd)
}

func UniverseAdd(args ...interface{}) {
	e := args[0].(*engine.Entity)
	l := e.Component(engine.COMP_Location).(*engine.Location)
	u := engine.Multiverse[l.UniverseId]

	if u == nil {
		util.Event.Fire(engine.EVNT_UniversePull, l.UniverseId)
		u = engine.Multiverse[l.UniverseId]
	}
	if u == nil {
		logrus.WithFields(logrus.Fields{
			"UniverseId": l.UniverseId,
		}).Warn("cmd.UniverseAdd: Failure: Universe not found")
		return
	}

	e.Id = u.GenerateId()
	u.Entities[e.Id] = e

	logrus.WithFields(logrus.Fields{
		"UniverseId": u.Id,
		"EntityId":   e.Id,
	}).Info("cmd.UniverseAdd")
}
