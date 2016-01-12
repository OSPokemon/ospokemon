package db

import (
	"github.com/ospokemon/ospokemon/server"
)

func ChangePassword(account *server.Account) error {
	_, err := Connection.Exec("UPDATE accounts SET password=? WHERE username=?", account.Password, account.Username)

	if err != nil {
		return err
	}

	return nil
}
