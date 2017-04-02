package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func init() {
	event.On(event.AccountsInsert, AccountsInsertPlayer)
}

func AccountsInsertPlayer(args ...interface{}) {
	account := args[0].(*ospokemon.Account)
	player, _ := ospokemon.GetPlayer(account.Username)

	if player == nil {
		class, _ := ospokemon.GetClass(0)
		entity := ospokemon.MakeEntity()
		player = ospokemon.BuildPlayer(account.Username, ospokemon.DEFAULT_BAG_SIZE, class, entity)
		player.Username = account.Username
		log.Add("Username", player.Username).Debug("account insert player: grant empty player")

		if err := ospokemon.Players.Insert(player); err != nil {
			log.Add("Username", player.Username).Add("Error", err.Error()).Error("Account insert player")
		}
	}
}
