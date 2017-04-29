package server

import (
	"ospokemon.com/log"
	"ospokemon.com/server/routes/logout"
	"ospokemon.com/server/sessionman"
	"time"
)

func PollSessionExpirations() {
	for now := range time.Tick(1 * time.Second) {
		for _, session := range sessionman.Cache {
			if session.Expire.Before(now) {

				log.Add("Username", session.Username).Add("SessionId", session.SessionId).Info("session expired")
				logout.LogoutPlayer(session.Username)

				if session.Websocket != nil {
					session.Websocket.Close()
				}
			}
		}
	}
}
