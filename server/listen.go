package server

import (
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/option"
	"ospokemon.com/server/api/logout"
	"ospokemon.com/server/session"
)

func Listen(s *session.Session) {
	for s.Websocket != nil {
		if message, err := s.Receive(); err == nil {
			go ReceiveMessage(s, message)
		} else {
			s.Websocket.Close()

			if err.Error() != "EOF" {
				log.Add("Error", err).Error("websocket error")
			}

			account := ospokemon.Accounts.Cache[s.Username]
			if account == nil {
				return
			}

			log.Add("Username", s.Username).Add("Universe", account.GetEntity().UniverseId).Add("SessionId", s.SessionId).Info("websocket closed")

			if !option.Bool("allow-refresh") {
				logout.LogoutPlayer(s.Username)
			} else {
				logout.RemoveEntity(s.Username)
			}

			return
		}
	}
}
