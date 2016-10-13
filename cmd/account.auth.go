package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
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
		if account, err := save.AccountsGet(username); err == nil {
			a = account
			save.Accounts[username] = a
		} else {
			logrus.WithFields(logrus.Fields{
				"Username": username,
			}).Warn("cmd.AccountAuth: Failure: Username not found")

			http.Redirect(w, r, "/login/?usernamenotfound", http.StatusMovedPermanently)
			return
		}
	}

	if a.Password != password {
		logrus.WithFields(logrus.Fields{
			"Username": a.Username,
			"Email":    a.Email,
		}).Warn("cmd.AccountAuth: Failure: Incorrect password")

		http.Redirect(w, r, "/login/?passwordwrong", http.StatusMovedPermanently)
		return
	}

	s := server.Sessions[a.SessionId]

	if s == nil {
		s = server.NewSession(a.Username)
		a.SessionId = s.SessionId
		server.Sessions[s.SessionId] = s
	}

	util.Event.Fire(save.EVNT_AccountLogin, a, s, r, w)

	s.WriteSessionId(w)
	http.Redirect(w, r, "/play/", http.StatusMovedPermanently)

	logrus.WithFields(logrus.Fields{
		"Username":  a.Username,
		"SessionId": a.SessionId,
	}).Warn("cmd.AccountAuth")
}
