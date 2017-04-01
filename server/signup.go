package server

import (
	"net/http"
	"ospokemon.com"
	"ospokemon.com/query"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	account := ospokemon.MakeAccount(r.FormValue("username"))
	account.Password = hashpassword(r.FormValue("password"))

	if err := query.AccountsInsert(account); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/login/#"+account.Username, http.StatusMovedPermanently)
})
