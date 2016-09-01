package server

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	a := &save.Account{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: hashpassword(r.FormValue("password")),
	}

	util.Log.WithFields(map[string]interface{}{
		"Username": a.Username,
		"Email":    a.Email,
	}).Warn("ospokemon/server/Signup:")

	util.Event.Fire(save.EVNT_AccountCreate, a, w)
})
