package server

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"time"
)

func PollSessionExpire() {
	for now := range time.Tick(1 * time.Second) {
		for _, s := range Sessions {
			if s.Expire.Before(now) {
				account := game.Accounts[s.Username]
				query.AccountsDelete(account)
				query.AccountsInsert(account)
				delete(game.Accounts, s.Username)
				delete(Sessions, s.SessionId)

				if s.Websocket != nil {
					s.Websocket.Close()
				}
			}
		}
	}
}
