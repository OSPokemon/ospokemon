package save

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Account struct {
	Username  string
	Password  string
	SessionId uint
	Register  time.Time
	part.Parts
}

func MakeAccount(username string) *Account {
	a := &Account{
		Username: username,
		Register: time.Now(),
		Parts:    make(part.Parts),
	}

	event.Fire(event.AccountMake, a)

	return a
}

var Accounts = make(map[string]*Account)

func GetAccount(username string) (*Account, error) {
	if a, ok := Accounts[username]; a != nil {
		return a, nil
	} else if ok {
		return nil, nil
	}

	a := MakeAccount(username)
	err := a.Query()
	if err != nil {
		a = nil
	}

	Accounts[username] = a
	return a, err
}

func (a *Account) Query() error {
	row := Connection.QueryRow(
		"SELECT password, register FROM accounts WHERE username=?",
		a.Username,
	)

	var timebuff int64
	if err := row.Scan(&a.Password, &timebuff); err != nil {
		return err
	}

	a.Register = time.Unix(timebuff, 0)

	event.Fire(event.AccountQuery, a)

	return nil
}

func (a *Account) Insert() error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, password, register) values (?, ?, ?)",
		a.Username,
		a.Password,
		a.Register.Unix(),
	)

	if err == nil {
		event.Fire(event.AccountInsert, a)
	}

	return err
}

func (a *Account) Delete() error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", a.Username)

	if err == nil {
		event.Fire(event.AccountDelete, a)
	}

	return err
}
