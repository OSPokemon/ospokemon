package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"net/http"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	if s := readsession(r); s != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	username := r.FormValue("username")
	password := hashpassword(r.FormValue("password"))

	if account := game.Accounts[username]; account != nil {
		if account.Password == password {
			session := account.Parts[PARTsession].(*Session)
			session.WriteSessionId(w)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		http.Redirect(w, r, "/login/?password#"+username, http.StatusMovedPermanently)
		return
	}

	account, err := query.GetAccount(username)

	if account == nil {
		logrus.WithFields(logrus.Fields{
			"Username": username,
		}).Debug("server.Login: account not found")

		http.Redirect(w, r, "/login/?account", http.StatusMovedPermanently)
		return
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": username,
			"Error":    err.Error(),
		}).Error("server.Login")

		http.Redirect(w, r, "/login/?account", http.StatusMovedPermanently)
		return
	}

	if account.Password == password {
		session := NewSession(username)
		session.WriteSessionId(w)
		Sessions[session.SessionId] = session

		entity := account.GetEntity()

		entity.AddPart(session)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, "/login/?password#"+username, http.StatusMovedPermanently)
	return
})
