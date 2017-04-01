package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.AccountsSelect, AccountsSelectPlayer)
}

func AccountsSelectPlayer(args ...interface{}) {
	account := args[0].(*ospokemon.Account)
	player, err := query.GetPlayer(account.Username)

	if player != nil {
		player.AddPart(account)
		player.AddPart(player)
		account.Parts = player.Parts
	} else if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("accounts select player")
	}
}
