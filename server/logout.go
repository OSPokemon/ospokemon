package server

import (
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		util.Event.Fire(EVNT_SessionExpire, s, r)
	}
})
