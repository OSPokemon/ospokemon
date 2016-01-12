package server

import (
	log "github.com/Sirupsen/logrus"
	"time"
)

type Account struct {
	Username    string
	Password    string
	Register    time.Time
	SessionId   int
	PlayerIds   []int
	PlayerId    int
	Permissions map[string]bool
}

var Accounts = make(map[string]*Account)

var LoadAccount func(username string) (*Account, error)
var CreateAccount func(username string, password string) (*Account, error)
var ChangePassword func(account *Account) error

func GetAccount(username string) *Account {
	if Accounts[username] == nil {
		if account, err := LoadAccount(username); err == nil {
			Accounts[username] = account
		} else {
			log.WithFields(log.Fields{
				"Username": username,
				"Error":    err.Error(),
			}).Info("Account lookup failed")
		}
	}

	return Accounts[username]
}
