package routes // import "ospokemon.com/server/routes"

import (
	"net/http"

	"ztaylor.me/http/mux"
)

func Routes(mux *mux.Mux, fs http.FileSystem, dbsalt string) {
	mux.MapLit("/ospokemon.js", OSPokemonJs)
	mux.MapLit("/api/signup", Signup(dbsalt))
	mux.MapLit("/api/login", Login(dbsalt))
	mux.MapLit("/api/logout", Logout)
	mux.MapLit("/api/websocket", Websocket)
	mux.MapLit("/", Page(fs))
}
