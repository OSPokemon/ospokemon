package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"time"
)

func AccountsSelect(username string) (*game.Account, error) {
	row := Connection.QueryRow(
		"SELECT password, register FROM accounts WHERE username=?",
		username,
	)

	account := game.MakeAccount(username)
	var timebuff int64
	err := row.Scan(&account.Password, &timebuff)

	if err != nil {
		game.Accounts[username] = nil
		return nil, err
	}

	account.Register = time.Unix(timebuff, 0)
	game.Accounts[username] = account

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Info("accounts select")

	event.Fire(event.AccountsSelect, account)
	return account, nil
}
