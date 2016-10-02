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
	log := logrus.WithFields(map[string]interface{}{
		"Username": username,
	})

	if err := playerdelete(username); err != nil {
		log.Error("cmd.PlayerDelete: " + err.Error())
	} else {
		log.Warn("cmd.PlayerDelete")
	}
}

func playerdelete(username string) error {
	_, err := save.Connection.Exec("DELETE FROM players WHERE username=?", username)
	return err
}
