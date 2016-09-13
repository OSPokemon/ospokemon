package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(engine.EVNT_UniverseRemove, UniverseRemove)
}

func UniverseRemove(args ...interface{}) {
	e := args[0].(*engine.Entity)
	entityid := e.Id
	l := e.Component(engine.COMP_Location).(*engine.Location)
	u := engine.Multiverse[l.UniverseId]

	if u == nil {
		logrus.WithFields(logrus.Fields{
			"UniverseId": l.UniverseId,
		}).Warn("cmd.UniverseRemove: Failure: Universe not found")
		return
	}

	delete(u.Entities, entityid)
	e.Id = 0

	logrus.WithFields(logrus.Fields{
		"UniverseId": u.Id,
		"EntityId":   entityid,
	}).Info("cmd.UniverseRemove")
}
