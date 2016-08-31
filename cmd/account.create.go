package cmd

import (
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"time"
)

func init() {
	util.Event.On(save.EVNT_AccountCreate, AccountCreate)
}

func AccountCreate(args ...interface{}) {
	a := args[0].(*save.Account)
	w := args[1].(http.ResponseWriter)

	_, err := save.Connection.Exec(`INSERT INTO accounts (
		username,
		email,
		password,
		register
	) values (?, ?, ?, ?)`,
		a.Username,
		a.Email,
		a.Password,
		time.Now().Unix(),
	)

	if err != nil {
		util.Log.Error(err)
		w.Write([]byte(err.Error()))
	} else {
		util.Event.Fire(save.EVNT_AccountLogin, a.Username, a.Password, w)
	}
}
