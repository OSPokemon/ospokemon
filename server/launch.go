package server

import (
	"net/http"

	"ztaylor.me/cast"
	"ztaylor.me/env"
)

func LaunchEnv(env env.Provider) {
	fs := http.Dir(env.Get("webpath"))
	dbsalt := env.Get("passwordsalt")
	server := New(fs, dbsalt)
	go PollSessionExpirations()

	if cast.Bool(env.Get("usehttps")) {
		server.ListenAndServeTLS("ospokemon.cert", "ospokemon.key")
	} else {
		server.ListenAndServe(":" + env.Get("port"))
	}
}
