package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

var fileserver = http.FileServer(http.Dir(util.Opt("webpath")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		s.Refresh()

		logrus.WithFields(logrus.Fields{
			"SessionId": s.SessionId,
			"Path":      r.RequestURI,
		}).Debug("server/PageHandler: Refresh Session")
	}

	fileserver.ServeHTTP(w, r)
})
