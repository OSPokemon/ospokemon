package server

import (
	"net/http"
	"ospokemon.com/log"
	"ospokemon.com/query"
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

	account, _ := query.GetAccount(username)
	if account == nil {
		log.Add("Username", username).Warn("login: account not found")
		http.Redirect(w, r, "/login/?account", http.StatusMovedPermanently)
		return
	}

	if account.Password != password {
		log.Add("Username", username).Warn("login: incorrect password")
		http.Redirect(w, r, "/login/?password#"+username, http.StatusMovedPermanently)
		return
	}

	session, _ := account.Parts[PARTsession].(*Session)
	if session == nil {
		session = NewSession(username)
		session.WriteSessionId(w)
		Sessions[session.SessionId] = session
		account.AddPart(session)
		log.Add("Username", username).Add("SessionId", session.SessionId).Info("login: create session")
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
})
