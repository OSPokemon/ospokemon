package server

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

var DoSignup func(string, string, int, int) error

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	username := r.FormValue("username")
	password := HashPassword(r.FormValue("password"))
	tc := r.FormValue("trainerclass")
	trainerclass, err := strconv.Atoi(tc)
	sid := r.FormValue("pokemon")
	speciesid, err := strconv.Atoi(sid)

	err = DoSignup(username, password, trainerclass, speciesid)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	log.WithFields(log.Fields{
		"Username":     username,
		"TrainerClass": trainerclass,
		"SpeciesId":    speciesid,
	}).Info("Account created")

	LoginHandler(w, r)
})
