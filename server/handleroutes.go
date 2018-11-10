package server

import (
	"net/http"

	"ospokemon.com/server/routes"
	"ztaylor.me/env"
)

func HandleRoutes(env env.Provider) {
	http.Handle("/ospokemon.js", routes.OSPokemonJs)
	http.Handle("/api/signup", routes.Signup)
	http.Handle("/api/login", routes.Login)
	http.Handle("/api/logout", routes.Logout)
	http.Handle("/api/websocket", routes.Websocket)
	http.Handle("/", routes.Page(env))
}
