package routes

import (
	"net/http"

	"ospokemon.com/server/sessionman"
	"ztaylor.me/env"
)

func Page(env env.Provider) http.Handler {
	fileserver := http.FileServer(http.Dir(env.Get("webpath")))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session := sessionman.FromRequestCookie(r); session != nil {
			session.Refresh()
		}

		fileserver.ServeHTTP(w, r)
	})
}
