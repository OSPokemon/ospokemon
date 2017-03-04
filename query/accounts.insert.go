package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
)

func AccountsInsert(account *game.Account) error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, password, register) values (?, ?, ?)",
		account.Username,
		account.Password,
		account.Register.Unix(),
	)

	if err == nil {
		delete(game.Accounts, account.Username)

		logrus.WithFields(logrus.Fields{
			"Username": account.Username,
		}).Info("accounts insert")

		event.Fire(event.AccountsInsert, account)
	}

	return nil
}
