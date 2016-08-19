package save

import (
	"time"
)

const EVNT_AccountCreate = "ospokemon/save/Account.Create"
const EVNT_AccountLogin = "ospokemon/save/Account.Login"

type Account struct {
	Username    string
	Email       string
	Password    string
	SessionId   uint
	Register    time.Time
	Permissions map[string]bool
}

var Accounts = make(map[string]*Account)
