package websocket

import (
	"ospokemon.com"
	"ospokemon.com/server/routes/logout"
	"ospokemon.com/server/sessionman"
	"ztaylor.me/cast"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

func Listen(session *sessionman.Session) {
	env := env.Global()
	for session.Websocket != nil {
		if message, err := session.Receive(); err == nil {
			go ReceiveMessage(session, message)
		} else {
			session.Websocket.Close()

			if err.Error() != "EOF" {
				log.Add("Error", err).Error("websocket error")
			}

			account := ospokemon.Accounts.Cache[session.Username]
			if account == nil {
				return
			}

			log.Add("Username", session.Username).Add("Universe", account.GetEntity().UniverseId).Add("SessionId", session.SessionId).Info("websocket closed")

			if !cast.Bool(env.Get("allow-refresh")) {
				logout.LogoutPlayer(session.Username)
			} else {
				logout.RemoveEntity(session.Username)
			}

			return
		}
	}
}
