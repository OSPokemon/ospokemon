package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
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
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("account insert player: grant empty player")
	}

	err := query.PlayersInsert(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("Account insert player")
	}
}
