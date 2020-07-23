package server

import (
	"time"

	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/routes/logout"
	"github.com/ospokemon/ospokemon/server/sessionman"
)

func PollSessionExpirations() {
	for now := range time.Tick(1 * time.Second) {
		for _, session := range sessionman.Cache {
			if session.Expire.Before(now) {

				ospokemon.LOG().Add("Username", session.Username).Add("SessionId", session.SessionId).Info("session expired")
				logout.LogoutPlayer(session.Username)

				if session.Websocket != nil {
					session.Websocket.Close()
				}
			}
		}
	}
}
