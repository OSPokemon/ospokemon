package server

import (
	"net/http"
)

var ChangePasswordHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	sessionId := readSessionId(r)
	session := Sessions[sessionId]

	if session == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"Error\":\"SessionId required\"}"))
		return
	}

	account := GetAccount(session.Username)

	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"Error\":\"Account does not exist\"}"))
	}

	password := HashPassword(r.FormValue("password"))
	account.Password = password

	err := ChangePassword(account)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"Error\":\"" + err.Error() + "\"}"))
	}

	LogoutHandler(w, r)
})
