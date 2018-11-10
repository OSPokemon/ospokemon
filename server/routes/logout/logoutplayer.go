package logout

import (
	"ospokemon.com"
	"ospokemon.com/server/sessionman"
	"ztaylor.me/log"
)

func LogoutPlayer(username string) {
	account := ospokemon.Accounts.Cache[username]
	if account == nil {
		log.Add("Username", username).Warn("logout: account missing")
		return
	}

	sessionman.Remove(account)

	ospokemon.Accounts.Delete(account)
	ospokemon.Accounts.Insert(account)

	RemoveEntity(username)

	delete(ospokemon.Accounts.Cache, username)
	delete(ospokemon.Players.Cache, username)
}
