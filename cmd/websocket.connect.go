package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketConnect, WebsocketConnect)
}

func WebsocketConnect(args ...interface{}) {
	s := args[0].(*server.Session)

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  s.Username,
	}).Warn("server.WebsocketConnect")
}
