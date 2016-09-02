package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	if session := readsession(r); session != nil {
		http.Redirect(w, r, "/play", http.StatusMovedPermanently)
		return
	}

	username := r.FormValue("username")
	password := hashpassword(r.FormValue("password"))

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Warn("ospokemon/server/Login:")

	util.Event.Fire(save.EVNT_AccountLogin, username, password, w)
})
