package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		logrus.WithFields(logrus.Fields{
			"Username":  s.Username,
			"SessionId": s.SessionId,
		}).Warn("ospokemon/server/Logout:")

		util.Event.Fire(save.EVNT_AccountLogout, s.Username)
	}
})
