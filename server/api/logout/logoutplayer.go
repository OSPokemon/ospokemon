package logout

import (
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/server/session"
)

func LogoutPlayer(username string) {
	account := ospokemon.Accounts.Cache[username]
	if account == nil {
		log.Add("Username", username).Warn("logout: account missing")
		return
	}

	session.Remove(account)

	ospokemon.Accounts.Delete(account)
	ospokemon.Accounts.Insert(account)

	if player, _ := ospokemon.GetPlayer(username); player != nil {
		entity := player.GetEntity()
		universe := ospokemon.Multiverse[entity.UniverseId]

		universe.Remove(entity)
	}

	delete(ospokemon.Accounts.Cache, username)
	delete(ospokemon.Players.Cache, username)
}
