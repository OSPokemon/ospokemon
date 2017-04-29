package routes

import (
	ws "golang.org/x/net/websocket"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/server/routes/websocket"
	"ospokemon.com/server/sessionman"
)

var Websocket = ws.Handler(func(conn *ws.Conn) {
	session := sessionman.FromRequestCookie(conn.Request())
	if session == nil {
		log.Add("RemoteAddr", conn.Request().RemoteAddr).Debug("websocket: session error")
		return
	}

	session.Lock()
	session.Websocket = conn

	if p, err := ospokemon.GetPlayer(session.Username); err != nil {
		log.Add("Username", session.Username).Add("Error", err.Error()).Error("websocket: player error")
	} else if u, err := ospokemon.GetUniverse(p.GetEntity().UniverseId); err != nil {
		log.Add("Universe", p.GetEntity().UniverseId).Add("Error", err.Error()).Error("websocket: universe error")
	} else {
		log.Add("Universe", p.GetEntity().UniverseId).Add("Username", session.Username).Add("SessionId", session.SessionId).Info("websocket opened")

		u.Add(p.GetEntity())
		websocket.Listen(session)
	}

	session.Websocket = nil
	session.Unlock()
})
