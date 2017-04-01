package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/server"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*game.Account)
		session := account.Parts[server.PARTsession].(*server.Session)
		delete(server.Sessions, session.SessionId)
	})
}
