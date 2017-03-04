package server

import (
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		account := game.Accounts[s.Username]
		query.AccountsDelete(account)
		query.AccountsInsert(account)

		http.Redirect(w, r, "/login/#"+s.Username, http.StatusMovedPermanently)
	}
})
