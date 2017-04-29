package websocket

import (
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/option"
	"ospokemon.com/server/routes/logout"
	"ospokemon.com/server/sessionman"
)

func Listen(session *sessionman.Session) {
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

			if !option.Bool("allow-refresh") {
				logout.LogoutPlayer(session.Username)
			} else {
				logout.RemoveEntity(session.Username)
			}

			return
		}
	}
}
