package save

import (
	"time"
)

const EVNT_AccountCreate = "save/Account.Create"
const EVNT_AccountAuth = "save/Account.Auth"
const EVNT_AccountLogin = "save/Account.Login"
const EVNT_AccountLogout = "save/Account.Logout"

type Account struct {
	Username    string
	Email       string
	Password    string
	SessionId   uint
	Register    time.Time
	Permissions map[string]bool
}

var Accounts = make(map[string]*Account)
