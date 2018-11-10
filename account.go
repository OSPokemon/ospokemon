package ospokemon

import "time"

const PARTaccount = "account"

type Account struct {
	Username string
	Password string
	Register time.Time
	Parts
}

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
	if Accounts.Cache[username] == nil {
		if a, err := Accounts.Select(username); a != nil {
			Accounts.Cache[username] = a
		} else {
			return nil, err
		}
	}

	return Accounts.Cache[username], nil
}

// persistence headers
var Accounts = struct {
	Cache  map[string]*Account
	Select func(string) (*Account, error)
	Insert func(*Account) error
	Delete func(*Account) error
}{make(map[string]*Account), nil, nil, nil}
