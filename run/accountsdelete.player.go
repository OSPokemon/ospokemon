package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.AccountsDelete, func(args ...interface{}) {
		account := args[0].(*game.Account)
		player := account.GetPlayer()
		err := query.PlayersDelete(player)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Username": account.Username,
				"Error":    err.Error(),
			}).Error("Accounts delete player")
		}
	})
}
