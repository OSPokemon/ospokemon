package routes

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/routes/websocket"
	"github.com/ospokemon/ospokemon/server/sessionman"
	ws "golang.org/x/net/websocket"
)

var Websocket = ws.Handler(func(conn *ws.Conn) {
	session := sessionman.FromRequestCookie(conn.Request())
	if session == nil {
		ospokemon.LOG().Add("RemoteAddr", conn.Request().RemoteAddr).Debug("websocket: session error")
		return
	}

	session.Lock()
	session.Websocket = conn

	if p, err := ospokemon.GetPlayer(session.Username); err != nil {
		ospokemon.LOG().Add("Username", session.Username).Add("Error", err.Error()).Error("websocket: player error")
	} else if u, err := ospokemon.GetUniverse(p.GetEntity().UniverseId); err != nil {
		ospokemon.LOG().Add("Universe", p.GetEntity().UniverseId).Add("Error", err.Error()).Error("websocket: universe error")
	} else {
		ospokemon.LOG().Add("Universe", p.GetEntity().UniverseId).Add("Username", session.Username).Add("SessionId", session.SessionId).Info("websocket opened")

		u.Add(p.GetEntity())
		websocket.Listen(session)
	}

	session.Websocket = nil
	session.Unlock()
})
