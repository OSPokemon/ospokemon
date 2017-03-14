package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetAccount(username string) (*game.Account, error) {
	if game.Accounts[username] == nil {
		if a, err := AccountsSelect(username); a != nil {
			game.Accounts[username] = a
		} else {
			return nil, err
		}
	}

	return game.Accounts[username], nil
}
