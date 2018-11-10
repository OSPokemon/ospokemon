package routes

import (
	"net/http"

	"ospokemon.com"
	"ospokemon.com/server/routes/signup"
	"ospokemon.com/server/security"
)

var Signup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	account := ospokemon.MakeAccount(r.FormValue("username"))
	account.Password = security.HashPassword(r.FormValue("password"))

	signup.MakePlayer(account)

	if err := ospokemon.Accounts.Insert(account); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/login/#"+account.Username, 307)
})
