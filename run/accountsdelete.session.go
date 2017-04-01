package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/server"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*ospokemon.Account)
		session := account.Parts[server.PARTsession].(*server.Session)
		delete(server.Sessions, session.SessionId)
	})
}
