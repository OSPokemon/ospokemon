package server

import (
	"net/http"
	"ospokemon.com/log"
	"ospokemon.com/option"
)

func Launch() {
	routes()
	go PollSessionExpirations()

	if option.Bool("usehttps") {
		log.Error(http.ListenAndServeTLS(":443", "ospokemon.cert", "ospokemon.key", nil))
	} else {
		log.Error(http.ListenAndServe(":"+option.String("port"), nil))
	}
}
