package server

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		util.Log.WithFields(map[string]interface{}{
			"Username":  s.Username,
			"SessionId": s.SessionId,
		}).Warn("ospokemon/server/Logout:")

		util.Event.Fire(save.EVNT_AccountLogout, s.Username)
	}
})
