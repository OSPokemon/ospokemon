package server

import (
	"net/http"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/server/session"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	if s := session.Find(r); s != nil {
		http.Redirect(w, r, "/", 307)
		return
	}

	username := r.FormValue("username")
	password := hashpassword(r.FormValue("password"))

	account, _ := ospokemon.GetAccount(username)
	if account == nil {
		log.Add("Username", username).Warn("login: account not found")
		http.Redirect(w, r, "/login/?account", 307)
		return
	}

	if account.Password != password {
		log.Add("Username", username).Warn("login: incorrect password")
		delete(ospokemon.Accounts.Cache, username)
		delete(ospokemon.Players.Cache, username)
		http.Redirect(w, r, "/login/?password#"+username, 307)
		return
	}

	if s := session.Get(account); s == nil {
		s = session.Add(account)
		log.Add("Username", username).Add("SessionId", s.SessionId).Info("login: create session")
	}

	session.Get(account).WriteSessionId(w)
	http.Redirect(w, r, "/", 307)
})
