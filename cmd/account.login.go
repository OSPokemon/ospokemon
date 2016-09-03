package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"time"
)

func init() {
	util.Event.On(save.EVNT_AccountLogin, AccountLogin)
}

func AccountLogin(args ...interface{}) {
	a := args[0].(save.Account)
	s := args[1].(*server.Session)
	r := args[2].(*http.Request)
	w := args[3].(http.ResponseWriter)
}
