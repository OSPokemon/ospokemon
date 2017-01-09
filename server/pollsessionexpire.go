package server

import (
	"github.com/ospokemon/ospokemon/save"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range Sessions {
			if s.Expire.Before(now) {
				save.Accounts[s.Username].Delete()
				save.Accounts[s.Username].Insert()
				delete(save.Accounts, s.Username)
				delete(Sessions, s.SessionId)

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
