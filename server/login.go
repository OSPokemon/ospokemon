package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"net/http"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	if s := readsession(r); s != nil {
		http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
		return
	}

	username := r.FormValue("username")
	password := hashpassword(r.FormValue("password"))

	if account := save.Accounts[username]; account != nil {
		if account.Password == password {
			session := Sessions[account.SessionId]
			session.WriteSessionId(w)
			http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
			return
		}

		http.Redirect(w, r, "/login/?password", http.StatusMovedPermanently)
		return
	}

	account := save.MakeAccount(username)

	if err := account.Query(); err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": username,
		}).Error("server.Login: " + err.Error())

		http.Redirect(w, r, "/login/?account", http.StatusMovedPermanently)
		return
	}

	if account.Password == password {
		session := NewSession(username)
		account.SessionId = session.SessionId
		session.WriteSessionId(w)
		save.Accounts[username] = account
		Sessions[session.SessionId] = session

		http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, "/login/?password", http.StatusMovedPermanently)
	return
})
