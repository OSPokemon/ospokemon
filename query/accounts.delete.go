package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func AccountsDelete(account *game.Account) error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", account.Username)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": account.Username,
		}).Info("accounts delete")

		event.Fire(event.AccountsDelete, account)
	}

	return err
}
