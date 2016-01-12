package server

import (
	"net/http"
	"strconv"
)

var ChooseProfileHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	sessionId := readSessionId(r)

	if sessionId < 1 {
		w.WriteHeader(500)
		w.Write([]byte("Session required"))
		return
	}

	playerId, err := strconv.ParseInt(r.FormValue("PlayerId"), 10, 0)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	session := Sessions[sessionId]
	account := Accounts[session.Username]

	validPlayerId := false
	for _, id2 := range account.PlayerIds {
		validPlayerId = validPlayerId || int(playerId) == id2
	}

	if !validPlayerId {
		w.WriteHeader(500)
		w.Write([]byte("Cannot control that profile"))
	}

	account.PlayerId = int(playerId)
	http.Redirect(w, r, "/play", http.StatusMovedPermanently)
})
