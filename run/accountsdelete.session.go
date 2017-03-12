package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/server"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*game.Account)
		session := account.Parts[server.PARTsession].(*server.Session)
		delete(server.Sessions, session.SessionId)
	})
}
