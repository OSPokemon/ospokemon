package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/server"
)

func LoadAccount(name string) {
	if server.Accounts[name] != nil {
		return
	}

	account := &server.Account{}

	row := Connection.QueryRow("SELECT id, name, password FROM players WHERE name=?", name)
	err := row.Scan(&account.PlayerId, &account.Username, &account.Password)
	if err != nil {
		log.Fatal(err)
	}

	server.Accounts[name] = account
}
