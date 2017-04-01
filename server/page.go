package server

import (
	"net/http"
	"ospokemon.com/option"
)

var fileserver = http.FileServer(http.Dir(option.String("webpath")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		s.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
