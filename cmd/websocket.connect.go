package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketConnect, WebsocketConnect)
}

func WebsocketConnect(args ...interface{}) {
	s := args[0].(*server.Session)

	p := save.Players[s.Username]

	if p == nil {
		util.Event.Fire(save.EVNT_PlayerPull, s.Username)
		p = save.Players[s.Username]
	}
	if p == nil {
		logrus.WithFields(logrus.Fields{
			"SessionId": s.SessionId,
			"Username":  s.Username,
		}).Error("cmd.WebsocketConnect: Player not found")

		util.Event.Fire(server.EVNT_WebsocketDisconnect, s)
		return
	}

	logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  p.Username,
	}).Info("cmd.WebsocketConnect")

	l := p.Entity.Component(engine.COMP_Location).(*engine.Location)
	util.Event.Fire(engine.EVNT_UniverseAdd, l.UniverseId, p.Entity)
}
