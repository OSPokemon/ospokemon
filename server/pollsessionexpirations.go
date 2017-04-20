package server

import (
	"ospokemon.com/log"
	"ospokemon.com/server/api/logout"
	"ospokemon.com/server/session"
	"time"
)

func PollSessionExpirations() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range session.Sessions {
			if s.Expire.Before(now) {

				log.Add("Username", s.Username).Add("SessionId", s.SessionId).Info("session expired")
				logout.LogoutPlayer(s.Username)

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
