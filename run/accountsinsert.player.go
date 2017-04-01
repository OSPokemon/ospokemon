package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/query"
)

func init() {
	event.On(event.AccountsInsert, AccountsInsertPlayer)
}

func AccountsInsertPlayer(args ...interface{}) {
	account := args[0].(*ospokemon.Account)
	player := ospokemon.Players[account.Username]

	if player == nil {
		class, _ := query.GetClass(0)
		entity := ospokemon.MakeEntity()
		player = ospokemon.BuildPlayer(account.Username, ospokemon.DEFAULT_BAG_SIZE, class, entity)
		player.Username = account.Username
		log.Add("Username", player.Username).Debug("account insert player: grant empty player")
	}

	err := query.PlayersInsert(player)

	if err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("Account insert player")
	}
}
