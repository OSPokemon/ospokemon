package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"time"
)

func init() {
	util.Event.On(save.EVNT_AccountCreate, AccountCreate)
}

func AccountCreate(args ...interface{}) {
	a := args[0].(*save.Account)
	r := args[1].(*http.Request)
	w := args[2].(http.ResponseWriter)

	// TODO: check for existing account

	_, err := save.Connection.Exec(
		"INSERT INTO accounts (username, email, password, register) values (?, ?, ?, ?)",
		a.Username,
		a.Email,
		a.Password,
		time.Now().Unix(),
	)

	if err != nil {
		logrus.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	logrus.WithFields(map[string]interface{}{
		"Username": a.Username,
		"Email":    a.Email,
	}).Warn("cmd/AccountCreate")

	util.Event.Fire(save.EVNT_AccountAuth, a.Username, a.Password, r, w)
}
