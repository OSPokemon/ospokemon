package server

import (
	"ospokemon.com/server/api/logout"
	"ospokemon.com/server/session"
	"time"
)

func PollSessionExpirations() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range session.Sessions {
			if s.Expire.Before(now) {

				logout.LogoutPlayer(s.Username)

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}