package server

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"net/http"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	account := game.MakeAccount(r.FormValue("username"))
	account.Password = hashpassword(r.FormValue("password"))

	if err := query.AccountsInsert(account); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/login/#"+account.Username, http.StatusMovedPermanently)
})
