package loader

import (
	"errors"
	"github.com/ospokemon/ospokemon/server"
)

func LoginAccount(username string, password string) (*server.Account, error) {
	if server.Accounts[username] == nil {
		account := &server.Account{}

		row := Connection.QueryRow("SELECT id, name, password FROM players WHERE name=?", username)
		err := row.Scan(&account.PlayerId, &account.Username, &account.Password)
		if err != nil {
			return nil, err
		}

		server.Accounts[username] = account
	}

	if server.Accounts[username].Password != password {
		return nil, errors.New("Password mismatch")
	}

	return server.Accounts[username], nil
}
