package db

import (
	"github.com/ospokemon/ospokemon/server"
	"time"
)

func CreateAccount(username string, password string) (*server.Account, error) {
	account := &server.Account{
		Username: username,
		Password: password,
		Register: time.Now(),
	}

	_, err := Connection.Exec("INSERT INTO accounts (username, password, register) VALUES (?, ?, ?)", account.Username, account.Password, account.Register.Unix())

	if err != nil {
		return nil, err
	}

	return account, nil
}
