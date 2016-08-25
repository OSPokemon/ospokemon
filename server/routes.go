package server

import (
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

func routes() {
	http.Handle("/", http.FileServer(http.Dir(util.Opt("webpath"))))
	http.Handle("/api/signup", SignupHandler)
	http.Handle("/api/login", LoginHandler)
	http.Handle("/api/logout", LogoutHandler)
}