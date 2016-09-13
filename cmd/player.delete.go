package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(save.EVNT_PlayerDelete, PlayerDelete)
}

func PlayerDelete(args ...interface{}) {
	username := args[0].(string)

	_, err := save.Connection.Exec("DELETE FROM players WHERE username=?", username)

	if err != nil {
		logrus.Error("cmd.PlayerDelete: " + err.Error())
		return
	}

	logrus.WithFields(map[string]interface{}{
		"Username": username,
	}).Warn("cmd.PlayerDelete")
}
