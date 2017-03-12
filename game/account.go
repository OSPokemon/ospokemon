package game

import (
	"time"
)

const PARTaccount = "account"

type Account struct {
	Username string
	Password string
	Register time.Time
	Parts
}

var Accounts = make(map[string]*Account)

func MakeAccount(username string) *Account {
	a := &Account{
		Username: username,
		Parts:    make(Parts),
	}

	return a
}

func (a *Account) Part() string {
	return PARTaccount
}

func (parts Parts) GetAccount() *Account {
	account, _ := parts[PARTaccount].(*Account)
	return account
}
