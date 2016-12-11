package server

import (
	"github.com/ospokemon/ospokemon/save"
	"net/http"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	account := save.MakeAccount(r.FormValue("username"))
	account.Password = hashpassword(r.FormValue("password"))

	if err := account.Insert(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
})
