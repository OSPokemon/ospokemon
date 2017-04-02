package ospokemon

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

var accounts = make(map[string]*Account)

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

func GetAccount(username string) (*Account, error) {
	if accounts[username] == nil {
		if a, err := Accounts.Select(username); a != nil {
			accounts[username] = a
		} else {
			return nil, err
		}
	}

	return accounts[username], nil
}

// persistence headers
var Accounts struct {
	Select func(string) (*Account, error)
	Insert func(*Account) error
	Delete func(*Account) error
}
