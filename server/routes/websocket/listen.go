package websocket

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/routes/logout"
	"github.com/ospokemon/ospokemon/server/sessionman"
	"taylz.io/types"
)

func Listen(session *sessionman.Session) {
	env := ospokemon.ENV()
	for session.Websocket != nil {
		if message, err := session.Receive(); err == nil {
			go ReceiveMessage(session, message)
		} else {
			session.Websocket.Close()

			if err.Error() != "EOF" {
				ospokemon.LOG().Add("Error", err).Error("websocket error")
			}

			account := ospokemon.Accounts.Cache[session.Username]
			if account == nil {
				return
			}

			ospokemon.LOG().Add("Username", session.Username).Add("Universe", account.GetEntity().UniverseId).Add("SessionId", session.SessionId).Info("websocket closed")

			if !types.BoolString(env["allow-refresh"]) {
				logout.LogoutPlayer(session.Username)
			} else {
				logout.RemoveEntity(session.Username)
			}

			return
		}
	}
}
