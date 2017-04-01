package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func AccountsInsert(account *ospokemon.Account) error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, password, register) values (?, ?, ?)",
		account.Username,
		account.Password,
		account.Register.Unix(),
	)

	if err == nil {
		delete(ospokemon.Accounts, account.Username)

		log.Add("Username", account.Username).Info("accounts insert")

		event.Fire(event.AccountsInsert, account)
	}

	return nil
}
