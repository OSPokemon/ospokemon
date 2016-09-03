package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
)

func Launch() {
	routes()
	go PollSessionExpire()
	e := http.ListenAndServe(":"+util.Opt("port"), nil)
	logrus.Error(e)
}
