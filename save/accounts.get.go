package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

func AccountsGet(username string) (*Account, error) {
	a := &Account{}
	row := Connection.QueryRow(
		"SELECT username, email, password, register FROM accounts WHERE username=?",
		username,
	)

	var timebuff int64
	if err := row.Scan(&a.Username, &a.Email, &a.Password, &timebuff); err == nil {
		a.Register = time.Unix(timebuff, 0)
	} else {
		return nil, errors.New("accountsget: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": a.Username,
	}).Debug("save.AccountsGet")

	return a, nil
}
