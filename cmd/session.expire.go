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
	s := args[0].(*server.Session)

	if s.Websocket != nil {
		util.Event.Fire(server.EVNT_WebsocketDisconnect, s)
	}

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  s.Username,
	}).Warn("server.SessionExpire")

	delete(server.Sessions, s.SessionId)
	s.SessionId = 0

	util.Event.Fire(save.EVNT_AccountLogout, s.Username)
}
