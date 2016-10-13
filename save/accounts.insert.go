package save

import (
	"errors"
)

var AccountsInsert = func(a *Account) error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, email, password, register) values (?, ?, ?, ?)",
		a.Username,
		a.Email,
		a.Password,
		a.Register.Unix(),
	)

	if err != nil {
		return errors.New("accountsinsert: " + err.Error())
	}

	return nil
}
