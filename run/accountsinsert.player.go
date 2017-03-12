package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.AccountsInsert, AccountsInsertPlayer)
}

func AccountsInsertPlayer(args ...interface{}) {
	account := args[0].(*game.Account)
	player := game.Players[account.Username]

	if player == nil {
		class, _ := query.GetClass(0)
		entity := game.MakeEntity()
		player = game.BuildPlayer(account.Username, game.DEFAULT_BAG_SIZE, class, entity)
		player.Username = account.Username
		log.Add("Username", "2").Debug("account insert player: grant empty player")
	}

	err := query.PlayersInsert(player)

	if err != nil {
		log.Add("Username", "2").Add("Error", err.Error()).Error("Account insert player")
	}
}
