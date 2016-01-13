package server

import (
	"net/http"
	"strconv"
)

var ChooseTrainerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	TrainerId, err := strconv.ParseInt(r.FormValue("TrainerId"), 10, 0)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	session := Sessions[sessionId]
	account := Accounts[session.Username]

	validTrainerId := false
	for _, id2 := range account.TrainerIds {
		validTrainerId = validTrainerId || int(TrainerId) == id2
	}

	if !validTrainerId {
		w.WriteHeader(500)
		w.Write([]byte("Cannot control that profile"))
	}

	account.TrainerId = int(TrainerId)
	http.Redirect(w, r, "/play", http.StatusMovedPermanently)
})
