package server

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	if sessionId := readsession(r); sessionId != nil {
		logrus.WithFields(logrus.Fields{
			"SessionId": sessionId,
		}).Warn("server/LoginHandler: Redirect session login")

		http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
		return
	}

	username := r.FormValue("username")
	password := hashpassword(r.FormValue("password"))

	util.Event.Fire(save.EVNT_AccountAuth, username, password, r, w)
})
