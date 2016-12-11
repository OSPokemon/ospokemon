package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"golang.org/x/net/websocket"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	s := readsession(conn.Request())

	if s == nil {
		logrus.Warn("server.WebsocketHandler: Failure: Session missing")
		return
	}

	s.Lock()
	defer s.Unlock()

	s.Websocket = conn

	p, err := save.GetPlayer(s.Username)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": s.Username,
		}).Error("server.Websocket: " + err.Error())
		return
	}

	p.Entity.AddComponent(s)

	location := p.Entity.Component(save.COMP_Location).(*save.Location)
	u, _ := save.GetUniverse(location.UniverseId)
	u.Add(p.Entity)

	Listen(s)
})
