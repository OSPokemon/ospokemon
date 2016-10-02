package server

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		a := save.Accounts[s.Username]
		util.Event.Fire(save.EVNT_AccountLogout, a)
	}
})
