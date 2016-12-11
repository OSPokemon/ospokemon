package server

import (
	"github.com/ospokemon/ospokemon/save"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range Sessions {
			if s.Expire.Before(now) {
				save.Accounts[s.Username].Update()
				save.Accounts[s.Username] = nil
				Sessions[s.SessionId] = nil

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
