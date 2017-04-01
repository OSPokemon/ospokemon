package server

import (
	"golang.org/x/net/websocket"
	"ospokemon.com/log"
	"ospokemon.com/query"
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
		log.Add("Username", s.Username).Add("Error", err.Error()).Error("WebsocketHandler")
		return
	}

	p.AddPart(s)

	e := p.GetEntity()

	if u, err := query.GetUniverse(e.UniverseId); err != nil {
		log.Add("Universe", e.UniverseId).Add("Error", err.Error()).Error("WebsocketHandler")
	} else {
		u.Add(e)
		Listen(s)
	}
})
