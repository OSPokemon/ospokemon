package server // import "github.com/ospokemon/ospokemon/server"

import (
	"net/http"

	"github.com/ospokemon/ospokemon/server/routes"
	"taylz.io/http/server"
)

func New(fs http.FileSystem, dbsalt string) *server.Mux {
	mux := &server.Mux{}
	routes.Routes(mux, fs, dbsalt)
	return mux
}
