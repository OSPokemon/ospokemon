package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func AccountsDelete(account *game.Account) error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", account.Username)

	if err == nil {
		log.Add("Username", account.Username).Info("accounts delete")

		event.Fire(event.AccountsDelete, account)
	}

	return err
}
