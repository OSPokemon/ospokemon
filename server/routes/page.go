package routes

import (
	"net/http"

	"ospokemon.com/server/sessionman"
)

func Page(fs http.FileSystem) http.Handler {
	fileserver := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session := sessionman.FromRequestCookie(r); session != nil {
			session.Refresh()
		}

		fileserver.ServeHTTP(w, r)
	})
}
