package server

import (
	"github.com/ospokemon/ospokemon/option"
	"net/http"
)

var fileserver = http.FileServer(http.Dir(option.String("webpath")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		s.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
