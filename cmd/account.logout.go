package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(save.EVNT_AccountLogout, AccountLogout)
}

func AccountLogout(args ...interface{}) {
	username := args[0].(string)

	delete(save.Accounts, username)

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Warn("cmd/AccountLogout")
}
