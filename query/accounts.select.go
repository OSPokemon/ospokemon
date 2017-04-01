package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"time"
)

func AccountsSelect(username string) (*ospokemon.Account, error) {
	row := Connection.QueryRow(
		"SELECT password, register FROM accounts WHERE username=?",
		username,
	)

	account := ospokemon.MakeAccount(username)
	var timebuff int64
	err := row.Scan(&account.Password, &timebuff)

	if err != nil {
		ospokemon.Accounts[username] = nil
		return nil, err
	}

	account.Register = time.Unix(timebuff, 0)

	log.Add("Username", username).Info("accounts select")

	event.Fire(event.AccountsSelect, account)
	return account, nil
}
