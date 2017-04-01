package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*game.Account)
		player := account.GetPlayer()
		err := query.PlayersDelete(player)

		if err != nil {
			log.Add("Username", player.Username).Add("Error", err.Error()).Error("Accounts delete player")
		}
	})
}
