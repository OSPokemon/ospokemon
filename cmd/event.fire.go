package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(util.EVNT_EventFire, EventFire)
}

func EventFire(args ...interface{}) {
	event := args[0].(string)

	logrus.Debug(event)
}
