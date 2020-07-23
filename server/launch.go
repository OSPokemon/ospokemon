package server

import (
	"net/http"

	"taylz.io/env"
	"taylz.io/types"
)

func LaunchEnv(env env.Service) {
	fs := http.Dir(env["webpath"])
	dbsalt := env["passwordsalt"]
	server := New(fs, dbsalt)
	go PollSessionExpirations()

	port := ":" + env["port"]

	if types.BoolString(env["usehttps"]) {
		http.ListenAndServeTLS(port, "ospokemon.cert", "ospokemon.key", server)
	} else {
		http.ListenAndServe(port, server)
	}
}
