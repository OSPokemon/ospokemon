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
	err := args[1].(string)

	if s.Websocket == nil {
		return
	}

	if err != "EOF" && err != server.EVNT_SessionExpire {
		logrus.WithFields(logrus.Fields{
			"SessionId": s.SessionId,
		}).Error(err)
	}

	s.Websocket.Close()
	s.Websocket = nil

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
	}).Warn("server.WebsocketDisconnect")
}
