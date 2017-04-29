package routes

import (
	"net/http"
	"ospokemon.com/option"
	"ospokemon.com/server/sessionman"
)

var fileserver = http.FileServer(http.Dir(option.String("webpath")))

var Page = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if session := sessionman.FromRequestCookie(r); session != nil {
		session.Refresh()
	}

	fileserver.ServeHTTP(w, r)
})
