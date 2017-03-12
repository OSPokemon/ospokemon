package server

import (
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/option"
	"net/http"
)

func Launch() {
	routes()
	go PollSessionExpire()
	e := http.ListenAndServe(":"+option.String("port"), nil)
	log.Error(e.Error())
}
