package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.AccountsSelect, AccountsSelectPlayer)
}

func AccountsSelectPlayer(args ...interface{}) {
	account := args[0].(*game.Account)
	player, err := query.GetPlayer(account.Username)

	if player != nil {
		player.AddPart(account)
		player.AddPart(player)
		account.Parts = player.Parts
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": account.Username,
			"Error":    err.Error(),
		}).Error("accounts select player")
	}
}
