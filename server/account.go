package server

var Accounts = make(map[string]*Account)

type Account struct {
	PlayerId  int
	Username  string
	Password  string
	SessionId int
}

var LoginAccount func(username string, password string) (*Account, error)

func LogoutAccount(account *Account) {
	if account.SessionId > 0 {
		delete(Sessions, account.SessionId)
	}
	delete(Accounts, account.Username)
}
