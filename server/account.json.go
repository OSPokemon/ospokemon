package server

import (
	"encoding/json"
	"net/http"
)

var AccountHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sessionId := readSessionId(r)

	if sessionId < 1 {
		w.WriteHeader(500)
		w.Write([]byte("Session missing"))
	}

	session := Sessions[sessionId]
	account := Accounts[session.Username]

	resp := make(map[string]interface{})
	resp["profiles"] = account.PlayerIds

	accountJson, _ := json.Marshal(resp)

	w.Write(accountJson)
})
