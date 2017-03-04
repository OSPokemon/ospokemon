package game

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Account struct {
	Username string
	Password string
	Register time.Time
	part.Parts
}

var Accounts = make(map[string]*Account)

func MakeAccount(username string) *Account {
	a := &Account{
		Username: username,
		Parts:    make(part.Parts),
	}

	return a
}

func (a *Account) Part() string {
	return part.Account
}
