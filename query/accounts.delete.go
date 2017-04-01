package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
)

func AccountsDelete(account *ospokemon.Account) error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", account.Username)

	if err == nil {
		log.Add("Username", account.Username).Info("accounts delete")

		event.Fire(event.AccountsDelete, account)
	}

	return err
}
