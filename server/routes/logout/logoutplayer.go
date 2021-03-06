package logout

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/sessionman"
)

func LogoutPlayer(username string) {
	account := ospokemon.Accounts.Cache[username]
	if account == nil {
		ospokemon.LOG().Add("Username", username).Warn("logout: account missing")
		return
	}

	sessionman.Remove(account)

	ospokemon.Accounts.Delete(account)
	ospokemon.Accounts.Insert(account)

	RemoveEntity(username)

	delete(ospokemon.Accounts.Cache, username)
	delete(ospokemon.Players.Cache, username)
}
