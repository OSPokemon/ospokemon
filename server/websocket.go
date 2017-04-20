package server

import (
	"golang.org/x/net/websocket"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/server/session"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	s := session.Find(conn.Request())
	if s == nil {
		log.Add("RemoteAddr", conn.Request().RemoteAddr).Debug("websocket: session error")
		return
	}

	s.Lock()
	defer s.Unlock()

	s.Websocket = conn

	if p, err := ospokemon.GetPlayer(s.Username); err != nil {
		log.Add("Username", s.Username).Add("Error", err.Error()).Error("websocket: player error")
	} else if u, err := ospokemon.GetUniverse(p.GetEntity().UniverseId); err != nil {
		log.Add("Universe", p.GetEntity().UniverseId).Add("Error", err.Error()).Error("websocket: universe error")
	} else {
		log.Add("Universe", p.GetEntity().UniverseId).Add("Username", s.Username).Add("SessionId", s.SessionId).Info("websocket opened")

		p.AddPart(s)
		u.Add(p.GetEntity())
		Listen(s)
	}
})
