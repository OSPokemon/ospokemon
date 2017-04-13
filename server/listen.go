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
				log.Warn(err.Error())
			}

			account := ospokemon.Accounts.Cache[s.Username]
			if account == nil {
				log.Add("Username", s.Username).Add("SessionId", s.SessionId).Warn("websocket: close: account missing")
				return
			}

			if !option.Bool("allow-refresh") {
				logout.LogoutPlayer(s.Username)
			}

			return
		}
	}
}
