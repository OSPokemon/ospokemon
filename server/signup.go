package server

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	account, err := CreateAccount(username, password)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	Accounts[account.Username] = account

	log.WithFields(log.Fields{
		"Username": account.Username,
	}).Info("Account created")

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
