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

	account := &save.Account{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: hashpassword(r.FormValue("password")),
	}

	util.Log.WithFields(map[string]interface{}{
		"Username": account.Username,
		"Email":    account.Email,
	}).Warn("ospokemon/server/signup: " + save.EVNT_AccountCreate)

	util.Event.Fire(save.EVNT_AccountCreate, account, w)
})
