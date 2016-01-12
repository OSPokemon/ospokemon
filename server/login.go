package server

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	sessionId := readSessionId(r)
	if session := Sessions[sessionId]; session != nil {
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
		return
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))

	account := GetAccount(username)

	if account == nil {
		http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
		return
	} else if account.Password != password {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}

	session := CreateSession()
	session.Username = account.Username
	account.SessionId = session.SessionId
	Sessions[session.SessionId] = session

	log.WithFields(log.Fields{
		"Account": session.Username,
		"Session": session.SessionId,
	}).Info("Session created")

	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(session.SessionId))
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
})
