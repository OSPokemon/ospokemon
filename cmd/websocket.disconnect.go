package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketDisconnect, WebsocketDisconnect)
}

func WebsocketDisconnect(args ...interface{}) {
	s := args[0].(*server.Session)

	if p := save.Players[s.Username]; p != nil {
		util.Event.Fire(save.EVNT_PlayerDelete, p.Username)
		util.Event.Fire(save.EVNT_PlayerPush, p.Username)
		util.Event.Fire(engine.EVNT_UniverseRemove, p.Entity)
	}

	s.Websocket.Close()
	s.Websocket = nil

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
	}).Warn("server.WebsocketDisconnect")
}
