package query

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
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

		log.Add("Username", account.Username).Info("accounts insert")

		event.Fire(event.AccountsInsert, account)
	}

	return nil
}
