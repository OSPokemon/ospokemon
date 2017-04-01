package query

import (
	"ospokemon.com"
)

func GetAccount(username string) (*ospokemon.Account, error) {
	if ospokemon.Accounts[username] == nil {
		if a, err := AccountsSelect(username); a != nil {
			ospokemon.Accounts[username] = a
		} else {
			return nil, err
		}
	}

	return ospokemon.Accounts[username], nil
}
