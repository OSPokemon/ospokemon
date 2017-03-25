package server

import (
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/option"
	"net/http"
)

func Launch() {
	routes()
	go PollSessionExpire()

	if option.Bool("usehttps") {
		log.Error(http.ListenAndServeTLS(":443", "ospokemon.cert", "ospokemon.key", nil))
	} else {
		log.Error(http.ListenAndServe(":"+option.String("port"), nil))
	}
}
