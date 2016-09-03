package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"time"
)

func init() {
	util.Event.On(save.EVNT_AccountAuth, AccountAuth)
}

func AccountAuth(args ...interface{}) {
	username := args[0].(string)
	password := args[1].(string)
	r := args[2].(*http.Request)
	w := args[3].(http.ResponseWriter)

	a := save.Accounts[username]

	if a == nil {
		a = queryaccount(username)
	}
	if a == nil {
		logrus.WithFields(logrus.Fields{
			"Username": username,
		}).Warn("cmd/AccountAuth: Failure: Username not found")

		http.Redirect(w, r, "/login/?usernamenotfound", http.StatusMovedPermanently)
		return
	}

	if a.Password != password {
		logrus.WithFields(logrus.Fields{
			"Username": a.Username,
			"Email":    a.Email,
		}).Warn("cmd/AccountAuth: Failure: Incorrect password")

		http.Redirect(w, r, "/login/?passwordwrong", http.StatusMovedPermanently)
		return
	}

	s := server.Sessions[a.SessionId]

	if s == nil {
		s = server.NewSession(a.Username)
		a.SessionId = s.SessionId
		save.Accounts[a.Username] = a
		server.Sessions[s.SessionId] = s

		logrus.WithFields(logrus.Fields{
			"Username": a.Username,
			"Session":  a.SessionId,
		}).Warn("cmd/AccountAuth: Success")
	} else {
		logrus.WithFields(logrus.Fields{
			"Username": username,
			"Session":  a.SessionId,
		}).Warn("cmd/AccountAuth: Success: Rewrite session")
	}

	s.WriteSessionId(w)
	http.Redirect(w, r, "/play/", http.StatusMovedPermanently)

	util.Event.Fire(save.EVNT_AccountLogin, a, s, r, w)
}

func queryaccount(username string) *save.Account {
	a := &save.Account{}
	row := save.Connection.QueryRow(
		"SELECT username, email, password, register FROM accounts WHERE username=?",
		username,
	)

	var timebuff int64
	if err := row.Scan(&a.Username, &a.Email, &a.Password, &timebuff); err == nil {
		a.Register = time.Unix(timebuff, 0)
	} else {
		logrus.Error(err)
		a = nil
	}

	return a
}
