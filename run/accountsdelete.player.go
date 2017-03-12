package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*game.Account)
		player := account.GetPlayer()
		err := query.PlayersDelete(player)

		if err != nil {
			log.Add("Username", "2").Add("Error", err.Error()).Error("Accounts delete player")
		}
	})
}
