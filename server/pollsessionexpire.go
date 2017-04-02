package server

import (
	"ospokemon.com"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range Sessions {
			if s.Expire.Before(now) {
				account, _ := ospokemon.GetAccount(s.Username)

				if account != nil {
					ospokemon.Accounts.Delete(account)
					ospokemon.Accounts.Insert(account)
				}

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
