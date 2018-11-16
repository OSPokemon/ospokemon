package server // import "ospokemon.com/server"

import (
	"net/http"

	"ospokemon.com/server/routes"
	"ztaylor.me/http/mux"
)

func New(fs http.FileSystem, dbsalt string) *mux.Mux {
	mux := mux.NewMux()
	routes.Routes(mux, fs, dbsalt)
	return mux
}
