package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_SessionExpire, SessionExpire)
}

func SessionExpire(args ...interface{}) {
	sessionId := args[0].(uint)
	s := args[1].(*server.Session)
	s.SessionId = 0
	delete(server.Sessions, sessionId)

	logrus.WithFields(logrus.Fields{
		"SessionId": sessionId,
		"Username":  s.Username,
	}).Warn("server/SessionExpire")

	util.Event.Fire(save.EVNT_AccountLogout, s.Username)
}
