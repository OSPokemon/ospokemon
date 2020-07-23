package routes // import "github.com/ospokemon/ospokemon/server/routes"

import (
	"net/http"

	"taylz.io/http/router"
	"taylz.io/http/server"
)

func Routes(mux *server.Mux, fs http.FileSystem, dbsalt string) {
	mux.Route(router.Path("/ospokemon.js"), OSPokemonJs)
	mux.Route(router.Path("/api/signup"), Signup(dbsalt))
	mux.Route(router.Path("/api/login"), Login(dbsalt))
	mux.Route(router.Path("/api/logout"), Logout)
	mux.Route(router.Path("/api/websocket"), Websocket)
	mux.Route(router.Path("/"), Page(fs))
}
