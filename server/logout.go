package server

import (
	"github.com/ospokemon/ospokemon/save"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		a := save.Accounts[s.Username]
		a.Update()
		delete(save.Accounts, s.Username)
		delete(Sessions, s.SessionId)

		http.Redirect(w, r, "/login/", http.StatusMovedPermanently)
	}
})
