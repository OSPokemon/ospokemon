package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketDisconnect, WebsocketDisconnect)
}

func WebsocketDisconnect(args ...interface{}) {
	s := args[0].(*server.Session)

	s.Websocket.Close()
	s.Websocket = nil

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
	}).Warn("server.WebsocketDisconnect")
}
