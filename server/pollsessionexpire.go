package server

import (
	"ospokemon.com"
	"ospokemon.com/query"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range Sessions {
			if s.Expire.Before(now) {
				account := ospokemon.Accounts[s.Username]

				if account != nil {
					query.AccountsDelete(account)
					query.AccountsInsert(account)
				}

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
