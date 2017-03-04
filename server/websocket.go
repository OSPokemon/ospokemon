package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
	"golang.org/x/net/websocket"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	s := readsession(conn.Request())

	if s == nil {
		return
	}

	s.Lock()
	defer s.Unlock()

	s.Websocket = conn

	p, err := query.GetPlayer(s.Username)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": s.Username,
		}).Error("server.Websocket: " + err.Error())
		return
	}

	p.AddPart(s)

	e := p.Parts[part.Entity].(*game.Entity)

	if u, err := query.GetUniverse(e.UniverseId); err != nil {
		logrus.WithFields(logrus.Fields{
			"Universe": e.UniverseId,
		}).Error(err.Error())
	} else {
		u.Add(e)
		Listen(s)
	}
})
