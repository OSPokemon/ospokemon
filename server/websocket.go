package server

import (
	"golang.org/x/net/websocket"
	"ospokemon.com"
	"ospokemon.com/log"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	s := readsession(conn.Request())

	if s == nil {
		log.Add("RemoteAddr", conn.Request().RemoteAddr).Debug("websocket: session error")
		return
	}

	s.Lock()
	defer s.Unlock()

	s.Websocket = conn

	p, err := ospokemon.GetPlayer(s.Username)
	if err != nil {
		log.Add("Username", s.Username).Add("Error", err.Error()).Error("websocket: player error")
		return
	}

	p.AddPart(s)

	e := p.GetEntity()

	if u, err := ospokemon.GetUniverse(e.UniverseId); err != nil {
		log.Add("Universe", e.UniverseId).Add("Error", err.Error()).Error("websocket: universe error")
	} else {
		u.Add(e)
		Listen(s)
	}
})
