package query

import (
	"github.com/ospokemon/ospokemon/game"
)

func GetAccount(username string) (*game.Account, error) {
	if a, ok := game.Accounts[username]; ok {
		return a, nil
	}

	return AccountsSelect(username)
}
