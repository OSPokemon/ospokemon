package server

import (
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sessionId := readSessionId(r)

	if session := Sessions[sessionId]; session != nil {
		delete(Accounts, session.Username)
		delete(Sessions, sessionId)
	}

	w.Header().Set("Set-Cookie", "SessionId=0")
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
})
