package server

import (
	"net/http"
	"ospokemon.com"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	account := ospokemon.MakeAccount(r.FormValue("username"))
	account.Password = hashpassword(r.FormValue("password"))

	if err := ospokemon.Accounts.Insert(account); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/login/#"+account.Username, 307)
})
