package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/option"
	"net/http"
)

func Launch() {
	routes()
	go PollSessionExpire()
	e := http.ListenAndServe(":"+option.String("port"), nil)
	logrus.Error(e)
}
