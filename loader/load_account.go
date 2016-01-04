package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/server"
)

func LoadAccount(name string) *server.Account {
	if server.Accounts[name] == nil {

		account := &server.Account{}

		row := Connection.QueryRow("SELECT id, name, password FROM players WHERE name=?", name)
		err := row.Scan(&account.PlayerId, &account.Username, &account.Password)
		if err != nil {
			log.Debug("Account lookup miss: ", name)
			account = nil
		}

		server.Accounts[name] = account
	}

	return server.Accounts[name]
}
