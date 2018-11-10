package server

import (
	"net/http"

	"ztaylor.me/cast"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

func Launch(env env.Provider) {
	HandleRoutes(env)
	go PollSessionExpirations()

	if cast.Bool(env.Get("usehttps")) {
		log.Error(http.ListenAndServeTLS(":443", "ospokemon.cert", "ospokemon.key", nil))
	} else {
		log.Error(http.ListenAndServe(":"+env.Get("port"), nil))
	}
}
