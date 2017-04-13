package server

import (
	"net/http"
	"ospokemon.com/option"
	"ospokemon.com/server/session"
)

var fileserver = http.FileServer(http.Dir(option.String("webpath")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := session.Find(r); s != nil {
		s.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
