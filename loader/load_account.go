package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/auth"
	"github.com/ospokemon/ospokemon/registry"
)

func init() {
	registry.AccountLoader = LoadAccount
}

func LoadAccount(name string) {
	if registry.Accounts[name] != nil {
		return
	}

	account := &auth.Account{}

	row := Connection.QueryRow("SELECT id, name, password FROM players WHERE name=?", name)
	err := row.Scan(&account.PlayerId, &account.Username, &account.Password)
	if err != nil {
		log.Fatal(err)
	}

	registry.Accounts[name] = account
}
